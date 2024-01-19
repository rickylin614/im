package repository

import (
	"im/internal/models/po"
	"im/internal/models/request"

	"gorm.io/gorm"
)

//go:generate mockery --name IExampleRepository --structname MockExampleRepository --filename mock_example.go --output mock_repository --outpkg mock_repository --with-expecter
type IExampleRepository interface {
	Get(db *gorm.DB, cond *request.ExampleGet) (*po.Example, error)
	GetList(db *gorm.DB, cond *request.ExampleGetList) (*po.PageResult[*po.Example], error)
	Create(db *gorm.DB, data *po.Example) (id any, err error)
	Update(db *gorm.DB, data *po.Example) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewExampleRepository(in digIn) IExampleRepository {
	return exampleRepository{in: in}
}

type exampleRepository struct {
	in digIn
}

func (h exampleRepository) Get(db *gorm.DB, cond *request.ExampleGet) (*po.Example, error) {
	result := &po.Example{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h exampleRepository) GetList(db *gorm.DB, cond *request.ExampleGetList) (*po.PageResult[*po.Example], error) {
	result := &po.PageResult[*po.Example]{
		Page: cond.GetPager(),
		Data: make([]*po.Example, 0),
	}
	db = db.Model(po.Example{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h exampleRepository) Create(db *gorm.DB, data *po.Example) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (h exampleRepository) Update(db *gorm.DB, data *po.Example) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (h exampleRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(po.Example{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
