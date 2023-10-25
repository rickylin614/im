package service

import (
	"context"
	"fmt"

	"im/internal/models"
	"im/internal/models/req"
	"im/internal/pkg/consts"
	"im/internal/util/crypto"
	"im/internal/util/errs"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
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

	// 產token並記錄
	token = uuid.New()
	loginRecord := composeLoginRecord(ctx, user, consts.LoginStateSuccess)
	go s.loginRecord(loginRecord)
	if err = s.in.Repository.UsersRepo.SetToken(ctx, token, user); err != nil {
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

func (s usersService) Logout(ctx *gin.Context, token string) (err error) {
	return s.in.Repository.UsersRepo.DelToken(ctx, token)
}

func (s usersService) GetByToken(ctx *gin.Context, token string) (user *models.Users, err error) {
	return s.in.Repository.UsersRepo.GetByToken(ctx, token)
}
