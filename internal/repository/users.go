package repository

import (
	"im/internal/models"
	"im/internal/models/req"

	"gorm.io/gorm"
)

type IUsersRepository interface {
	Get(db *gorm.DB, cond *req.UsersGet) (*models.Users, error)
	GetList(db *gorm.DB, cond *req.UsersGetList) (*models.PageResult[*models.Users], error)
	Create(db *gorm.DB, data *models.Users) (id any, err error)
	Update(db *gorm.DB, data *models.Users) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewUsersRepository(in digIn) IUsersRepository {
	return usersRepository{in: in}
}

type usersRepository struct {
	in digIn
}

func (h usersRepository) Get(db *gorm.DB, cond *req.UsersGet) (*models.Users, error) {
	result := &models.Users{}
	if err := db.Find(result, cond).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h usersRepository) GetList(db *gorm.DB, cond *req.UsersGetList) (*models.PageResult[*models.Users], error) {
	result := &models.PageResult[*models.Users]{
		Page: cond.GetPager(),
		Data: make([]*models.Users, 0),
	}
	db = db.Model(models.Users{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond()).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (h usersRepository) Create(db *gorm.DB, data *models.Users) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (h usersRepository) Update(db *gorm.DB, data *models.Users) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (h usersRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(models.Users{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
