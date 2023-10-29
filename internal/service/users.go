package service

import (
	"context"
	"fmt"
	"time"

	"im/internal/models"
	"im/internal/models/req"
	"im/internal/pkg/consts"
	"im/internal/util/crypto"
	"im/internal/util/ctxs"
	"im/internal/util/errs"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	_ "github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
)

type IUsersService interface {
	Get(ctx *gin.Context, cond *req.UsersGet) (*models.Users, error)
	GetList(ctx *gin.Context, cond *req.UsersGetList) (*models.PageResult[*models.Users], error)
	Create(ctx *gin.Context, cond *req.UsersCreate) (id any, err error)
	Update(ctx *gin.Context, cond *req.UsersUpdate) (err error)
	Delete(ctx *gin.Context, cond *req.UsersDelete) (err error)
	Login(ctx *gin.Context, cond *req.UsersLogin) (token string, err error)
	Logout(ctx *gin.Context, token string) (err error)
	GetByToken(ctx *gin.Context, token string) (user *models.Users, err error)
}

func NewUsersService(in digIn) IUsersService {
	return usersService{in: in}
}

type usersService struct {
	in digIn
}

func (s usersService) Get(ctx *gin.Context, cond *req.UsersGet) (*models.Users, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.UsersRepo.Get(db, cond)
}

func (s usersService) GetList(ctx *gin.Context, cond *req.UsersGetList) (*models.PageResult[*models.Users], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.UsersRepo.GetList(db, cond)
}

func (s usersService) Create(ctx *gin.Context, cond *req.UsersCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.Users{ID: uuid.New(), PasswordHash: crypto.Hash(cond.Password)}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.UsersRepo.Create(db, insertData)
}

func (s usersService) Update(ctx *gin.Context, cond *req.UsersUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.Users{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.UsersRepo.Update(db, updateData)
}

func (s usersService) Delete(ctx *gin.Context, cond *req.UsersDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.UsersRepo.Delete(db, cond.ID)
}

func (s usersService) Login(ctx *gin.Context, cond *req.UsersLogin) (token string, err error) {
	db := s.in.DB.Session(ctx)

	// 取得使用者資訊
	getCond := &req.UsersGet{Username: cond.Username}
	user, err := s.in.Repository.UsersRepo.Get(db, getCond)
	if err != nil || user == nil {
		return "", errs.LoginCommonError
	}

	// 驗證密碼
	if user.PasswordHash != crypto.Hash(cond.Password) {
		loginRecord := composeLoginRecord(ctx, user, consts.LoginStateFailed)
		go s.loginRecord(loginRecord)
		return "", errs.LoginCommonError
	}

	// 驗證用戶狀態
	if user.Status != consts.UserStatusActive {
		loginRecord := composeLoginRecord(ctx, user, consts.LoginStateBlocked)
		go s.loginRecord(loginRecord)
		return "", errs.LoginLockedError
	}

	// 登入成功
	loginRecord := composeLoginRecord(ctx, user, consts.LoginStateSuccess)
	go s.loginRecord(loginRecord)

	// 產token並返回
	jwtTokenUuid := uuid.New()
	jwfclaims := jwt.NewWithClaims(jwt.SigningMethodRS256,
		&models.JWTClaims{
			User:     user,
			DeviceID: ctxs.GetDeviceID(ctx),
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    user.ID,
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(consts.TOKEN_EXPIRED)), // 假設token有效期為24小時
				ID:        jwtTokenUuid,
			},
		})
	token, err = jwfclaims.SignedString(crypto.GetRsaPrivateKey())
	if err != nil {
		s.in.Logger.Error(ctx, fmt.Errorf("sjwfclaims.SignedString err: %w", err))
		return "", errs.CommonServiceUnavailable
	}

	if err = s.in.Repository.UsersRepo.SetToken(ctx, user.ID, ctxs.GetDeviceID(ctx), token); err != nil {
		s.in.Logger.Error(ctx, fmt.Errorf("service set token err: %w", err))
		return "", errs.CommonServiceUnavailable
	}
	return token, nil
}

func (s usersService) loginRecord(loginRecord *models.LoginRecord) {
	ctx := context.Background()
	db := s.in.DB.Session(ctx)
	if _, err := s.in.Repository.LoginRecordRepo.Create(db, loginRecord); err != nil {
		s.in.Logger.Error(ctx, fmt.Errorf("login record create err: %w", err))
	}
}

func composeLoginRecord(ctx *gin.Context, user *models.Users, loginStatus consts.LoginState) *models.LoginRecord {
	return &models.LoginRecord{
		Name:       user.Nickname,
		UserID:     user.ID,
		UserAgent:  ctx.Request.UserAgent(),
		Ip:         ctx.ClientIP(),
		RemoteIp:   ctx.RemoteIP(),
		LoginState: loginStatus,
	}
}

func (s usersService) Logout(ctx *gin.Context, jwtToken string) (err error) {
	token, err := jwt.ParseWithClaims(jwtToken, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保token的签名算法是我们预期的
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return crypto.GetRsaPublicKey(), nil
	})

	if err != nil {
		s.in.Logger.Error(ctx, err)
		return errs.RequestTokenError
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok || token.Valid {
		return errs.RequestTokenError
	}

	return s.in.Repository.UsersRepo.DelToken(ctx, claims.Issuer, claims.DeviceID)
}

func (s usersService) GetByToken(ctx *gin.Context, jwtToken string) (user *models.Users, err error) {
	token, err := jwt.ParseWithClaims(jwtToken, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 驗證算法
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return crypto.GetRsaPublicKey(), nil
	})

	if err != nil {
		s.in.Logger.Error(ctx, err)
		return nil, errs.RequestTokenError
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok || !token.Valid {
		return nil, errs.RequestTokenError
	}
	_, err = s.in.Repository.UsersRepo.GetByToken(ctx, claims.Issuer, claims.DeviceID, claims.ID)
	if err != nil {
		return nil, err
	}

	return claims.User, nil
}
