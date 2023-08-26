package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

type ExampleRepository struct {
	in digIn
}

func (h ExampleRepository) ExampleGet(db *gorm.DB, cond req.ExampleGet) (*models.Example, error) {
	result := &models.Example{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h ExampleRepository) ExampleGetList(db *gorm.DB, cond req.ExampleGetList) (*models.PageResult[*models.Example], error) {
	result := &models.PageResult[*models.Example]{
		Page: cond.GetPager(),
		Data: make([]*models.Example, 0),
	}
	if err := db.Model(result).Find(result.Data, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h ExampleRepository) ExamplePost(db *gorm.DB, data *models.Example) (err error) {
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (h ExampleRepository) ExamplePut(db *gorm.DB, data *models.Example) (err error) {
	if err := db.Save(data).Error; err != nil {
		return err
	}
	return nil
}

func (h ExampleRepository) ExampleDelete(db *gorm.DB, data *models.Example) (err error) {
	if err := db.Delete(data).Error; err != nil {
		return err
	}
	return nil
}
