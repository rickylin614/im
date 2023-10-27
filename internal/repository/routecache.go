package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

type IRouteCacheRepository interface {
	Get(db *gorm.DB, cond *req.RouteCacheGet) (*models.RouteCache, error)
	GetList(db *gorm.DB, cond *req.RouteCacheGetList) (*models.PageResult[*models.RouteCache], error)
	Create(db *gorm.DB, data *models.RouteCache) (id any, err error)
	Update(db *gorm.DB, data *models.RouteCache) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewRouteCacheRepository(in digIn) IRouteCacheRepository {
	return routeCacheRepository{in: in}
}

type routeCacheRepository struct {
	in digIn
}

func (r routeCacheRepository) Get(db *gorm.DB, cond *req.RouteCacheGet) (*models.RouteCache, error) {
	result := &models.RouteCache{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r routeCacheRepository) GetList(db *gorm.DB, cond *req.RouteCacheGetList) (*models.PageResult[*models.RouteCache], error) {
	result := &models.PageResult[*models.RouteCache]{
		Page: cond.GetPager(),
		Data: make([]*models.RouteCache, 0),
	}
	db = db.Model(models.RouteCache{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r routeCacheRepository) Create(db *gorm.DB, data *models.RouteCache) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r routeCacheRepository) Update(db *gorm.DB, data *models.RouteCache) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r routeCacheRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.RouteCache{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
