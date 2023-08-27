package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

type IExampleRepository interface {
	Get(db *gorm.DB, cond *req.ExampleGet) (*models.Example, error)
	GetList(db *gorm.DB, cond *req.ExampleGetList) (*models.PageResult[*models.Example], error)
	Create(db *gorm.DB, data *models.Example) (err error)
	Update(db *gorm.DB, data *models.Example) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewExampleRepository(in digIn) IExampleRepository {
	return ExampleRepository{in: in}
}

type ExampleRepository struct {
	in digIn
}

func (h ExampleRepository) Get(db *gorm.DB, cond *req.ExampleGet) (*models.Example, error) {
	result := &models.Example{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h ExampleRepository) GetList(db *gorm.DB, cond *req.ExampleGetList) (*models.PageResult[*models.Example], error) {
	result := &models.PageResult[*models.Example]{
		Page: cond.GetPager(),
		Data: make([]*models.Example, 0),
	}
	if err := db.Model(result).Find(result.Data, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h ExampleRepository) Create(db *gorm.DB, data *models.Example) (err error) {
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (h ExampleRepository) Update(db *gorm.DB, data *models.Example) (err error) {
	if err := db.Save(data).Error; err != nil {
		return err
	}
	return nil
}

func (h ExampleRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.Example{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
