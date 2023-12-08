package repository

import (
	"{{ .ProjectName }}/internal/models"
	"{{ .ProjectName }}/internal/models/req"

	"gorm.io/gorm"
)

//go:generate mockery --name I{{ .FileName }}Repository --structname Mock{{ .FileName }}Repository --filename mock_{{ .FileNameSnakeCase }}.go --output mock_repository --outpkg mock_repository --with-expecter

type I{{ .FileName }}Repository interface {
	Get(db *gorm.DB, cond *req.{{ .FileName }}Get) (*models.{{ .FileName }}, error)
	GetList(db *gorm.DB, cond *req.{{ .FileName }}GetList) (*models.PageResult[*models.{{ .FileName }}], error)
	Create(db *gorm.DB, data *models.{{ .FileName }}) (id any, err error)
	Update(db *gorm.DB, data *models.{{ .FileName }}) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func New{{ .FileName }}Repository(in digIn) I{{ .FileName }}Repository {
	return &{{ .FileNameTitleLower }}Repository{in: in}
}

type {{ .FileNameTitleLower }}Repository struct {
	in digIn
}

func (r *{{ .FileNameTitleLower }}Repository) Get(db *gorm.DB, cond *req.{{ .FileName }}Get) (*models.{{ .FileName }}, error) {
	result := &models.{{ .FileName }}{}
	db = db.Find(result, cond)
	if db.Error != nil {
		return nil, db.Error
	}
	if db.RowsAffected == 0 {
		return nil, nil
	}
	return result, nil
}

func (r *{{ .FileNameTitleLower }}Repository) GetList(db *gorm.DB, cond *req.{{ .FileName }}GetList) (*models.PageResult[*models.{{ .FileName }}], error) {
	result := &models.PageResult[*models.{{ .FileName }}]{
		Page: cond.GetPager(),
		Data: make([]*models.{{ .FileName }}, 0),
	}
	db = db.Model(models.{{ .FileName }}{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *{{ .FileNameTitleLower }}Repository) Create(db *gorm.DB, data *models.{{ .FileName }}) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r *{{ .FileNameTitleLower }}Repository) Update(db *gorm.DB, data *models.{{ .FileName }}) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r *{{ .FileNameTitleLower }}Repository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.{{ .FileName }}{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
