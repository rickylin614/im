package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

type IFriendRequestsRepository interface {
	Get(db *gorm.DB, cond *req.FriendRequestsGet) (*models.FriendRequests, error)
	GetList(db *gorm.DB, cond *req.FriendRequestsGetList) (*models.PageResult[*models.FriendRequests], error)
	Create(db *gorm.DB, data *models.FriendRequests) (id any, err error)
	Update(db *gorm.DB, data *models.FriendRequests) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewFriendRequestsRepository(in digIn) IFriendRequestsRepository {
	return FriendRequestsRepository{in: in}
}

type FriendRequestsRepository struct {
	in digIn
}

func (r FriendRequestsRepository) Get(db *gorm.DB, cond *req.FriendRequestsGet) (*models.FriendRequests, error) {
	result := &models.FriendRequests{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r FriendRequestsRepository) GetList(db *gorm.DB, cond *req.FriendRequestsGetList) (*models.PageResult[*models.FriendRequests], error) {
	result := &models.PageResult[*models.FriendRequests]{
		Page: cond.GetPager(),
		Data: make([]*models.FriendRequests, 0),
	}
	db = db.Model(models.FriendRequests{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond()).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r FriendRequestsRepository) Create(db *gorm.DB, data *models.FriendRequests) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r FriendRequestsRepository) Update(db *gorm.DB, data *models.FriendRequests) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r FriendRequestsRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.FriendRequests{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
