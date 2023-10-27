package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

type IGroupMembersRepository interface {
	Get(db *gorm.DB, cond *req.GroupMembersGet) (*models.GroupMembers, error)
	GetList(db *gorm.DB, cond *req.GroupMembersGetList) (*models.PageResult[*models.GroupMembers], error)
	Create(db *gorm.DB, data *models.GroupMembers) (id any, err error)
	Update(db *gorm.DB, data *models.GroupMembers) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewGroupMembersRepository(in digIn) IGroupMembersRepository {
	return groupMembersRepository{in: in}
}

type groupMembersRepository struct {
	in digIn
}

func (r groupMembersRepository) Get(db *gorm.DB, cond *req.GroupMembersGet) (*models.GroupMembers, error) {
	result := &models.GroupMembers{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r groupMembersRepository) GetList(db *gorm.DB, cond *req.GroupMembersGetList) (*models.PageResult[*models.GroupMembers], error) {
	result := &models.PageResult[*models.GroupMembers]{
		Page: cond.GetPager(),
		Data: make([]*models.GroupMembers, 0),
	}
	db = db.Model(models.GroupMembers{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r groupMembersRepository) Create(db *gorm.DB, data *models.GroupMembers) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.GroupID, nil
}

func (r groupMembersRepository) Update(db *gorm.DB, data *models.GroupMembers) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r groupMembersRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.GroupMembers{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
