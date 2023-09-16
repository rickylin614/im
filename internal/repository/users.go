package repository

import (
	"context"
	"encoding/json"
	"im/internal/consts"
	"im/internal/models"
	"im/internal/models/req"
	"time"

	"gorm.io/gorm"
)

type IUsersRepository interface {
	Get(db *gorm.DB, cond *req.UsersGet) (*models.Users, error)
	GetList(db *gorm.DB, cond *req.UsersGetList) (*models.PageResult[*models.Users], error)
	Create(db *gorm.DB, data *models.Users) (id any, err error)
	Update(db *gorm.DB, data *models.Users) (err error)
	Delete(db *gorm.DB, id string) (err error)
	GetByToken(ctx context.Context, token string) (*models.Users, error)
	SetToken(ctx context.Context, token string, user *models.Users) error
	DelToken(ctx context.Context, token string) error
}

func NewUsersRepository(in digIn) IUsersRepository {
	return usersRepository{in: in}
}

type usersRepository struct {
	in digIn
}

func (r usersRepository) Get(db *gorm.DB, cond *req.UsersGet) (*models.Users, error) {
	result := &models.Users{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r usersRepository) GetList(db *gorm.DB, cond *req.UsersGetList) (*models.PageResult[*models.Users], error) {
	result := &models.PageResult[*models.Users]{
		Page: cond.GetPager(),
		Data: make([]*models.Users, 0),
	}
	db = db.Model(models.Users{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond()).Find(&result.Data).Error; err != nil {
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

func (r usersRepository) GetByToken(ctx context.Context, token string) (*models.Users, error) {
	result := &models.Users{}
	key := consts.LoginKey + token
	rdata, err := r.in.Rdb.Get(ctx, key).Bytes()
	if err != nil {
		r.in.Logger.Info(ctx, err.Error())
		return nil, err
	}
	if err := json.Unmarshal(rdata, result); err != nil {
		r.in.Logger.Error(ctx, err)
		return nil, err
	}
	r.in.Rdb.Expire(ctx, token, time.Hour*2)
	return result, nil
}

func (r usersRepository) SetToken(ctx context.Context, token string, user *models.Users) error {
	key := consts.LoginKey + token
	data, _ := json.Marshal(user)
	return r.in.Rdb.Set(ctx, key, data, time.Hour*2).Err()
}

func (r usersRepository) DelToken(ctx context.Context, token string) error {
	key := consts.LoginKey + token
	return r.in.Rdb.Del(ctx, key).Err()
}
