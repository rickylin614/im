package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

type ILoginRecordRepository interface {
	Get(db *gorm.DB, cond *req.LoginRecordGet) (*models.LoginRecord, error)
	GetList(db *gorm.DB, cond *req.LoginRecordGetList) (*models.PageResult[*models.LoginRecord], error)
	Create(db *gorm.DB, data *models.LoginRecord) (id any, err error)
	Update(db *gorm.DB, data *models.LoginRecord) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewLoginRecordRepository(in digIn) ILoginRecordRepository {
	return loginRecordRepository{in: in}
}

type loginRecordRepository struct {
	in digIn
}

func (h loginRecordRepository) Get(db *gorm.DB, cond *req.LoginRecordGet) (*models.LoginRecord, error) {
	result := &models.LoginRecord{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h loginRecordRepository) GetList(db *gorm.DB, cond *req.LoginRecordGetList) (*models.PageResult[*models.LoginRecord], error) {
	result := &models.PageResult[*models.LoginRecord]{
		Page: cond.GetPager(),
		Data: make([]*models.LoginRecord, 0),
	}
	db = db.Model(models.LoginRecord{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond()).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h loginRecordRepository) Create(db *gorm.DB, data *models.LoginRecord) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (h loginRecordRepository) Update(db *gorm.DB, data *models.LoginRecord) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (h loginRecordRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.LoginRecord{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
