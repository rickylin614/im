package repository

import (
	"im/internal/models/po"
	"im/internal/models/request"

	"gorm.io/gorm"
)

//go:generate mockery --name IMessageRepository --structname MockMessageRepository --filename mock_message.go --output mock_repository --outpkg mock_repository --with-expecter

type IMessageRepository interface {
	Get(db *gorm.DB, cond *request.MessageGet) (*po.Message, error)
	GetList(db *gorm.DB, cond *request.MessageGetList) (*po.PageResult[*po.Message], error)
	Create(db *gorm.DB, data *po.Message) (id any, err error)
	Update(db *gorm.DB, data *po.Message) (err error)
	Delete(db *gorm.DB, id string) (err error)
}

func NewMessageRepository(in digIn) IMessageRepository {
	return &messageRepository{in: in}
}

type messageRepository struct {
	in digIn
}

func (r *messageRepository) Get(db *gorm.DB, cond *request.MessageGet) (*po.Message, error) {
	result := &po.Message{}
	db = db.Find(result, cond)
	if db.Error != nil {
		return nil, db.Error
	}
	if db.RowsAffected == 0 {
		return nil, nil
	}
	return result, nil
}

func (r *messageRepository) GetList(db *gorm.DB, cond *request.MessageGetList) (*po.PageResult[*po.Message], error) {
	result := &po.PageResult[*po.Message]{
		Page: cond.GetPager(),
		Data: make([]*po.Message, 0),
	}
	db = db.Model(po.Message{}).Scopes(cond.Scope)
	if err := db.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(result.PagerCond).Find(&result.Data).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *messageRepository) Create(db *gorm.DB, data *po.Message) (id any, err error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data.ID, nil
}

func (r *messageRepository) Update(db *gorm.DB, data *po.Message) (err error) {
	if err := db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r *messageRepository) Delete(db *gorm.DB, id string) (err error) {
	if err := db.Model(po.Message{}).Delete("where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
