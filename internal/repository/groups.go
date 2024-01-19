package repository

import (
	"im/internal/models/po"
	"im/internal/models/request"

	"gorm.io/gorm"
)

//go:generate mockery --name IGroupsRepository --structname MockGroupsRepository --filename mock_groups.go --output mock_repository --outpkg mock_repository --with-expecter
type IGroupsRepository interface {
	Get(db *gorm.DB, cond *request.GroupsGet) (*po.Groups, error)
	GetList(db *gorm.DB, cond *request.GroupsGetList) (*po.PageResult[*po.Groups], error)
	Create(db *gorm.DB, data *po.Groups) (id any, err error)
	Update(db *gorm.DB, data *po.Groups) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewGroupsRepository(in digIn) IGroupsRepository {
	return groupsRepository{in: in}
}

type groupsRepository struct {
	in digIn
}

func (r groupsRepository) Get(db *gorm.DB, cond *request.GroupsGet) (*po.Groups, error) {
	result := &po.Groups{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r groupsRepository) GetList(db *gorm.DB, cond *request.GroupsGetList) (*po.PageResult[*po.Groups], error) {
	result := &po.PageResult[*po.Groups]{
		Page: cond.GetPager(),
		Data: make([]*po.Groups, 0),
	}
	db = db.Model(po.Groups{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r groupsRepository) Create(db *gorm.DB, data *po.Groups) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r groupsRepository) Update(db *gorm.DB, data *po.Groups) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r groupsRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(po.Groups{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
