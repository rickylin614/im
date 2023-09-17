package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

type IFriendRepository interface {
	Get(db *gorm.DB, cond *req.FriendGet) (*models.Friend, error)
	GetList(db *gorm.DB, cond *req.FriendGetList) (*models.PageResult[*models.Friend], error)
	Create(db *gorm.DB, data *models.Friend) (id any, err error)
	Update(db *gorm.DB, data *models.Friend) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewFriendRepository(in digIn) IFriendRepository {
	return friendRepository{in: in}
}

type friendRepository struct {
	in digIn
}

func (r friendRepository) Get(db *gorm.DB, cond *req.FriendGet) (*models.Friend, error) {
	result := &models.Friend{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r friendRepository) GetList(db *gorm.DB, cond *req.FriendGetList) (*models.PageResult[*models.Friend], error) {
	result := &models.PageResult[*models.Friend]{
		Page: cond.GetPager(),
		Data: make([]*models.Friend, 0),
	}
	db = db.Model(models.Friend{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond()).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r friendRepository) Create(db *gorm.DB, data *models.Friend) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r friendRepository) Update(db *gorm.DB, data *models.Friend) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r friendRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.Friend{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
