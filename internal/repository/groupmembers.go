package repository

import (
	"context"

	"im/internal/models/po"
	"im/internal/models/request"
	"im/internal/pkg/consts/rediskey"
	"im/internal/util/cache"

	"time"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

//go:generate mockery --name IGroupMembersRepository --structname MockGroupMembersRepository --filename mock_groupmembers.go --output mock_repository --outpkg mock_repository --with-expecter
type IGroupMembersRepository interface {
	Get(db *gorm.DB, cond *request.GroupMembersGet) (*po.GroupMembers, error)
	GetList(db *gorm.DB, cond *request.GroupMembersGetList) (*po.PageResult[*po.GroupMembers], error)
	GetListById(ctx context.Context, db *gorm.DB, cond *request.GroupMembersGetList) ([]*po.GroupMembers, error)
	Create(db *gorm.DB, data *po.GroupMembers) (id any, err error)
	Update(db *gorm.DB, data *po.GroupMembers) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewGroupMembersRepository(in digIn) IGroupMembersRepository {
	return &groupMembersRepository{in: in, sf: &singleflight.Group{}}
}

type groupMembersRepository struct {
	in digIn
	sf *singleflight.Group
}

// Get 確認成員資訊時使用
func (r *groupMembersRepository) Get(db *gorm.DB, cond *request.GroupMembersGet) (*po.GroupMembers, error) {
	result := &po.GroupMembers{}
	db = db.Find(result, cond)
	if db.Error != nil {
		return nil, db.Error
	}
	if db.RowsAffected == 0 {
		return nil, nil
	}
	return result, nil
}

func (r *groupMembersRepository) GetList(db *gorm.DB, cond *request.GroupMembersGetList) (*po.PageResult[*po.GroupMembers], error) {
	result := &po.PageResult[*po.GroupMembers]{
		Data: make([]*po.GroupMembers, 0),
	}
	db = db.Model(po.GroupMembers{}).Scopes(cond.Scope)
	if err := db.Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *groupMembersRepository) GetListById(ctx context.Context, db *gorm.DB, cond *request.GroupMembersGetList) ([]*po.GroupMembers, error) {
	key := rediskey.GROUP_MEMBER_KEY + cond.Id
	return cache.GetCache[[]*po.GroupMembers](ctx,
		r.in.Cache, r.in.Rdb, r.sf, key,
		func() ([]*po.GroupMembers, error) {
			result := make([]*po.GroupMembers, 0)
			db = db.Model(po.GroupMembers{}).Scopes(cond.Scope)
			if err := db.Find(&result).Error; err != nil {
				return nil, err
			}
			return result, nil
		}, time.Minute)
}

func (r *groupMembersRepository) Create(db *gorm.DB, data *po.GroupMembers) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.GroupID, nil
}

func (r *groupMembersRepository) Update(db *gorm.DB, data *po.GroupMembers) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r *groupMembersRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(po.GroupMembers{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
