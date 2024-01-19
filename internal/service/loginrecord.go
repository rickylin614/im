package service

import (
	"im/internal/models/po"
	"im/internal/models/request"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ILoginRecordService interface {
	Get(ctx *gin.Context, cond *request.LoginRecordGet) (*po.LoginRecord, error)
	GetList(ctx *gin.Context, cond *request.LoginRecordGetList) (*po.PageResult[*po.LoginRecord], error)
	Create(ctx *gin.Context, cond *request.LoginRecordCreate) (id any, err error)
	Update(ctx *gin.Context, cond *request.LoginRecordUpdate) (err error)
	Delete(ctx *gin.Context, cond *request.LoginRecordDelete) (err error)
}

func NewLoginRecordService(in DigIn) ILoginRecordService {
	return loginRecordService{in: in}
}

type loginRecordService struct {
	in DigIn
}

func (s loginRecordService) Get(ctx *gin.Context, cond *request.LoginRecordGet) (*po.LoginRecord, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.LoginRecordRepo.Get(db, cond)
}

func (s loginRecordService) GetList(ctx *gin.Context, cond *request.LoginRecordGetList) (*po.PageResult[*po.LoginRecord], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.LoginRecordRepo.GetList(db, cond)
}

func (s loginRecordService) Create(ctx *gin.Context, cond *request.LoginRecordCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &po.LoginRecord{}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.LoginRecordRepo.Create(db, insertData)
}

func (s loginRecordService) Update(ctx *gin.Context, cond *request.LoginRecordUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &po.LoginRecord{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.LoginRecordRepo.Update(db, updateData)
}

func (s loginRecordService) Delete(ctx *gin.Context, cond *request.LoginRecordDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.LoginRecordRepo.Delete(db, cond.ID)
}
