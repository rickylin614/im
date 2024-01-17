package repository

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"im/internal/models"
	"im/internal/models/request"
	"im/internal/pkg/consts"
	"im/internal/pkg/consts/rediskey"
	"im/internal/util/crypto"
	"im/internal/util/errs"
)

//go:generate mockery --name IUsersRepository --structname MockUsersRepository --filename mock_users.go --output mock_repository --outpkg mock_repository --with-expecter
type IUsersRepository interface {
	Get(db *gorm.DB, cond *request.UsersGet) (*models.Users, error)
	GetList(db *gorm.DB, cond *request.UsersGetList) (*models.PageResult[*models.Users], error)
	Create(db *gorm.DB, data *models.Users) (id any, err error)
	Update(db *gorm.DB, data *models.Users) (err error)
	Delete(db *gorm.DB, id string) (err error)
	GetByToken(ctx context.Context, UserID, deviceID, reqToken string) (*models.JWTClaims, error)
	SetToken(ctx context.Context, UserID, deviceID, jwtData string) error
	DelToken(ctx context.Context, UserID, deviceID string) error
}

func NewUsersRepository(in digIn) IUsersRepository {
	return usersRepository{in: in}
}

type usersRepository struct {
	in digIn
}

func (r usersRepository) Get(db *gorm.DB, cond *request.UsersGet) (*models.Users, error) {
	result := &models.Users{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r usersRepository) GetList(db *gorm.DB, cond *request.UsersGetList) (*models.PageResult[*models.Users], error) {
	result := &models.PageResult[*models.Users]{
		Page: cond.GetPager(),
		Data: make([]*models.Users, 0),
	}
	db = db.Model(models.Users{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r usersRepository) Create(db *gorm.DB, data *models.Users) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r usersRepository) Update(db *gorm.DB, data *models.Users) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r usersRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.Users{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r usersRepository) GetByToken(ctx context.Context, UserID, deviceID, reqToken string) (*models.JWTClaims, error) {
	key := rediskey.LoginKey + UserID + ":" + deviceID
	rget := r.in.Rdb.Get(ctx, key)
	if rget.Err() != nil {
		r.in.Logger.Error(ctx, rget.Err())
		return nil, rget.Err()
	}

	token, err := jwt.ParseWithClaims(rget.Val(), &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保token的签名算法是我们预期的
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return crypto.GetRsaPublicKey(), nil
	})

	if err != nil {
		r.in.Logger.Error(ctx, err)
		return nil, errs.RequestTokenError
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok || !token.Valid {
		r.in.Logger.Error(ctx, fmt.Errorf("token.Claims.(*models.JWTClaims) error, useId: %s, device: %s", UserID, deviceID))
		return nil, errs.RequestTokenError
	}

	if claims.ID != reqToken {
		r.in.Logger.Error(ctx, fmt.Errorf("token.Claims.(*models.JWTClaims) error, useId: %s, device: %s, claims.ID: %s, reqToken: %s",
			UserID, deviceID, claims.ID, reqToken))
		return nil, errs.RequestTokenError
	}

	// 延長token時效
	r.in.Rdb.Expire(ctx, key, consts.TOKEN_EXPIRED)
	return claims, nil
}

func (r usersRepository) SetToken(ctx context.Context, UserID, deviceID, jwtData string) error {
	key := rediskey.LoginKey + UserID + ":" + deviceID
	return r.in.Rdb.Set(ctx, key, jwtData, consts.TOKEN_EXPIRED).Err()
}

func (r usersRepository) DelToken(ctx context.Context, UserID, deviceID string) error {
	key := rediskey.LoginKey + UserID + ":" + deviceID
	return r.in.Rdb.Del(ctx, key).Err()
}
