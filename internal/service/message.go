package service

import (
	"im/internal/models/po"
	"im/internal/models/request"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IMessageService interface {
	Get(ctx *gin.Context, cond *request.MessageGet) (*po.Message, error)
	GetList(ctx *gin.Context, cond *request.MessageGetList) (*po.PageResult[*po.Message], error)
	Create(ctx *gin.Context, cond *request.MessageCreate) (id any, err error)
	Update(ctx *gin.Context, cond *request.MessageUpdate) (err error)
	Delete(ctx *gin.Context, cond *request.MessageDelete) (err error)
}

func NewMessageService(in DigIn) IMessageService {
	return messageService{In: in}
}

type messageService struct {
	In DigIn
}

func (s messageService) Get(ctx *gin.Context, cond *request.MessageGet) (*po.Message, error) {
	db := s.In.DB.Session(ctx)
	return s.In.Repository.MessageRepo.Get(db, cond)
}

func (s messageService) GetList(ctx *gin.Context, cond *request.MessageGetList) (*po.PageResult[*po.Message], error) {
	db := s.In.DB.Session(ctx)
	return s.In.Repository.MessageRepo.GetList(db, cond)
}

func (s messageService) Create(ctx *gin.Context, cond *request.MessageCreate) (id any, err error) {
	db := s.In.DB.Session(ctx)
	insertData := &po.Message{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.In.Repository.MessageRepo.Create(db, insertData)
}

func (s messageService) Update(ctx *gin.Context, cond *request.MessageUpdate) (err error) {
	db := s.In.DB.Session(ctx)
	updateData := &po.Message{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.In.Repository.MessageRepo.Update(db, updateData)
}

func (s messageService) Delete(ctx *gin.Context, cond *request.MessageDelete) (err error) {
	db := s.In.DB.Session(ctx)
	return s.In.Repository.MessageRepo.Delete(db, cond.ID)
}
