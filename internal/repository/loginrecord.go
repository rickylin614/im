package repository

import (
	"im/internal/models/po"
	"im/internal/models/request"

	"gorm.io/gorm"
)

//go:generate mockery --name ILoginRecordRepository --structname MockLoginRecordRepository --output mock_repository --outpkg mock_repository --filename mock_login_record.go --with-expecter
type ILoginRecordRepository interface {
	Get(db *gorm.DB, cond *request.LoginRecordGet) (*po.LoginRecord, error)
	GetList(db *gorm.DB, cond *request.LoginRecordGetList) (*po.PageResult[*po.LoginRecord], error)
	Create(db *gorm.DB, data *po.LoginRecord) (id any, err error)
	Update(db *gorm.DB, data *po.LoginRecord) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewLoginRecordRepository(in digIn) ILoginRecordRepository {
	return loginRecordRepository{in: in}
}

type loginRecordRepository struct {
	in digIn
}

func (r loginRecordRepository) Get(db *gorm.DB, cond *request.LoginRecordGet) (*po.LoginRecord, error) {
	result := &po.LoginRecord{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r loginRecordRepository) GetList(db *gorm.DB, cond *request.LoginRecordGetList) (*po.PageResult[*po.LoginRecord], error) {
	result := &po.PageResult[*po.LoginRecord]{
		Page: cond.GetPager(),
		Data: make([]*po.LoginRecord, 0),
	}
	db = db.Model(po.LoginRecord{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r loginRecordRepository) Create(db *gorm.DB, data *po.LoginRecord) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r loginRecordRepository) Update(db *gorm.DB, data *po.LoginRecord) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r loginRecordRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(po.LoginRecord{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
