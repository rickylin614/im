package repository

import (
	"im/internal/models/po"
	"im/internal/models/request"

	"gorm.io/gorm"
)

//go:generate mockery --name IFriendRequestsRepository --structname MockFriendRequestsRepository --filename mock_friendrequests.go --output mock_repository --outpkg mock_repository --with-expecter
type IFriendRequestsRepository interface {
	Get(db *gorm.DB, cond *request.FriendRequestsGet) (*po.FriendRequests, error)
	GetList(db *gorm.DB, cond *request.FriendRequestsGetList) (*po.PageResult[*po.FriendRequests], error)
	Create(db *gorm.DB, data *po.FriendRequests) (id any, err error)
	Update(db *gorm.DB, data *po.FriendRequests) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewFriendRequestsRepository(in digIn) IFriendRequestsRepository {
	return FriendRequestsRepository{in: in}
}

type FriendRequestsRepository struct {
	in digIn
}

func (r FriendRequestsRepository) Get(db *gorm.DB, cond *request.FriendRequestsGet) (*po.FriendRequests, error) {
	result := &po.FriendRequests{}
	if err := db.Scopes(cond.Scope).Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r FriendRequestsRepository) GetList(db *gorm.DB, cond *request.FriendRequestsGetList) (*po.PageResult[*po.FriendRequests], error) {
	result := &po.PageResult[*po.FriendRequests]{
		Page: cond.GetPager(),
		Data: make([]*po.FriendRequests, 0),
	}
	db = db.Model(po.FriendRequests{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r FriendRequestsRepository) Create(db *gorm.DB, data *po.FriendRequests) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r FriendRequestsRepository) Update(db *gorm.DB, data *po.FriendRequests) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r FriendRequestsRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(po.FriendRequests{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
