package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

//go:generate mockery --name IGroupInvitationRepository --structname MockGroupInvitationRepository --filename mock_group_invitation.go --output mock_repository --outpkg mock_repository --with-expecter

type IGroupInvitationRepository interface {
	Get(db *gorm.DB, cond *req.GroupInvitationGet) (*models.GroupInvitation, error)
	GetList(db *gorm.DB, cond *req.GroupInvitationGetList) (*models.PageResult[*models.GroupInvitation], error)
	Create(db *gorm.DB, data *models.GroupInvitation) (id any, err error)
	Update(db *gorm.DB, data *models.GroupInvitation) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewGroupInvitationRepository(in digIn) IGroupInvitationRepository {
	return groupInvitationRepository{in: in}
}

type groupInvitationRepository struct {
	in digIn
}

func (r groupInvitationRepository) Get(db *gorm.DB, cond *req.GroupInvitationGet) (*models.GroupInvitation, error) {
	result := &models.GroupInvitation{}
	db = db.Find(result, cond)
	if db.Error != nil {
		return nil, db.Error
	}
	if db.RowsAffected == 0 {
		return nil, nil
	}
	return result, nil
}

func (r groupInvitationRepository) GetList(db *gorm.DB, cond *req.GroupInvitationGetList) (*models.PageResult[*models.GroupInvitation], error) {
	result := &models.PageResult[*models.GroupInvitation]{
		Page: cond.GetPager(),
		Data: make([]*models.GroupInvitation, 0),
	}
	db = db.Model(models.GroupInvitation{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r groupInvitationRepository) Create(db *gorm.DB, data *models.GroupInvitation) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r groupInvitationRepository) Update(db *gorm.DB, data *models.GroupInvitation) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r groupInvitationRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.GroupInvitation{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
