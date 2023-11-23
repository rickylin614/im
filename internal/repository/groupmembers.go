package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

//go:generate mockery --name IGroupMembersRepository --structname MockGroupMembersRepository --filename mock_groupmembers.go --output mock_repository --outpkg mock_repository --with-expecter
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

// Get 確認成員資訊時使用
func (r groupMembersRepository) Get(db *gorm.DB, cond *req.GroupMembersGet) (*models.GroupMembers, error) {
	result := &models.GroupMembers{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r groupMembersRepository) GetList(db *gorm.DB, cond *req.GroupMembersGetList) (*models.PageResult[*models.GroupMembers], error) {
	result := &models.PageResult[*models.GroupMembers]{
		Data: make([]*models.GroupMembers, 0),
	}
	db = db.Model(models.GroupMembers{}).Scopes(cond.Scope)
	if err := db.Find(&result.Data).Error; err != nil {
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
