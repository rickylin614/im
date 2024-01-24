package repository

import (
	"{{ .ProjectName }}/internal/models/po"
	"{{ .ProjectName }}/internal/models/request"

	"gorm.io/gorm"
)

//go:generate mockery --name I{{ .FileName }}Repository --structname Mock{{ .FileName }}Repository --filename mock_{{ .FileNameSnakeCase }}.go --output mock_repository --outpkg mock_repository --with-expecter

type I{{ .FileName }}Repository interface {
	Get(db *gorm.DB, cond *request.{{ .FileName }}Get) (*po.{{ .FileName }}, error)
	GetList(db *gorm.DB, cond *request.{{ .FileName }}GetList) (*po.PageResult[*po.{{ .FileName }}], error)
	Create(db *gorm.DB, data *po.{{ .FileName }}) (id any, err error)
	Update(db *gorm.DB, data *po.{{ .FileName }}) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func New{{ .FileName }}Repository(in digIn) I{{ .FileName }}Repository {
	return &{{ .FileNameTitleLower }}Repository{in: in}
}

type {{ .FileNameTitleLower }}Repository struct {
	in digIn
}

func (r *{{ .FileNameTitleLower }}Repository) Get(db *gorm.DB, cond *request.{{ .FileName }}Get) (*po.{{ .FileName }}, error) {
	result := &po.{{ .FileName }}{}
	db = db.Find(result, cond)
	if db.Error != nil {
		return nil, db.Error
	}
	if db.RowsAffected == 0 {
		return nil, nil
	}
	return result, nil
}

func (r *{{ .FileNameTitleLower }}Repository) GetList(db *gorm.DB, cond *request.{{ .FileName }}GetList) (*po.PageResult[*po.{{ .FileName }}], error) {
	result := &po.PageResult[*po.{{ .FileName }}]{
		Page: cond.GetPager(),
		Data: make([]*po.{{ .FileName }}, 0),
	}
	db = db.Model(po.{{ .FileName }}{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *{{ .FileNameTitleLower }}Repository) Create(db *gorm.DB, data *po.{{ .FileName }}) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r *{{ .FileNameTitleLower }}Repository) Update(db *gorm.DB, data *po.{{ .FileName }}) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r *{{ .FileNameTitleLower }}Repository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(po.{{ .FileName }}{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
