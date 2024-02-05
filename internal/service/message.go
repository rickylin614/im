package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"im/internal/models/po"
	"im/internal/models/request"
)

type IMessageService interface {
	Get(ctx *gin.Context, cond *request.MessageGet) (*po.Message, error)
	GetList(ctx *gin.Context, cond *request.MessageGetList) (*po.PageResult[*po.Message], error)
	Create(ctx *gin.Context, cond *po.Message) (id any, err error)
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

func (s messageService) Create(ctx *gin.Context, insertData *po.Message) (id any, err error) {
	db := s.In.DB.Session(ctx)
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
