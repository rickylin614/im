package repository

import (
	"im/internal/models/po"
	"im/internal/models/request"

	"gorm.io/gorm"
)

//go:generate mockery --name IFriendRepository --structname MockFriendRepository --filename mock_friend.go --output mock_repository --outpkg mock_repository --with-expecter
type IFriendRepository interface {
	Get(db *gorm.DB, cond *request.FriendGet) (*po.Friend, error)
	GetList(db *gorm.DB, cond *request.FriendGetList) (*po.PageResult[*po.Friend], error)
	GetMutualList(db *gorm.DB, cond *request.FriendMutualGet) (*po.PageResult[*po.Friend], error)
	Create(db *gorm.DB, data *po.Friend) (id any, err error)
	Update(db *gorm.DB, data *po.Friend) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewFriendRepository(in digIn) IFriendRepository {
	return friendRepository{in: in}
}

type friendRepository struct {
	in digIn
}

func (r friendRepository) Get(db *gorm.DB, cond *request.FriendGet) (*po.Friend, error) {
	result := &po.Friend{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r friendRepository) GetList(db *gorm.DB, cond *request.FriendGetList) (*po.PageResult[*po.Friend], error) {
	result := &po.PageResult[*po.Friend]{
		Page: cond.GetPager(),
		Data: make([]*po.Friend, 0),
	}
	db = db.Model(po.Friend{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r friendRepository) GetMutualList(db *gorm.DB, cond *request.FriendMutualGet) (*po.PageResult[*po.Friend], error) {
	result := &po.PageResult[*po.Friend]{
		Page: cond.GetPager(),
		Data: make([]*po.Friend, 0),
	}
	// var commonFriends []Friend
	db = db.Table("friends").
		Joins("JOIN friends AS f2 ON friends.p_user_id = f2.p_user_id").
		Where("friends.f_user_id = ? AND f2.f_user_id = ?", cond.UserID, cond.TUserId)
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r friendRepository) Create(db *gorm.DB, data *po.Friend) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r friendRepository) Update(db *gorm.DB, data *po.Friend) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r friendRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(po.Friend{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
