package service

import (
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IFriendService interface {
	Get(ctx *gin.Context, cond *req.FriendGet) (*models.Friend, error)
	GetList(ctx *gin.Context, cond *req.FriendGetList) (*models.PageResult[*models.Friend], error)
	Create(ctx *gin.Context, cond *req.FriendCreate) (id any, err error)
	Update(ctx *gin.Context, cond *req.FriendUpdate) (err error)
	Delete(ctx *gin.Context, cond *req.FriendDelete) (err error)
}

func NewFriendService(in digIn) IFriendService {
	return friendService{in: in}
}

type friendService struct {
	in digIn
}

func (s friendService) Get(ctx *gin.Context, cond *req.FriendGet) (*models.Friend, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRepo.Get(db, cond)
}

func (s friendService) GetList(ctx *gin.Context, cond *req.FriendGetList) (*models.PageResult[*models.Friend], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRepo.GetList(db, cond)
}

func (s friendService) Create(ctx *gin.Context, cond *req.FriendCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.Friend{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.FriendRepo.Create(db, insertData)
}

func (s friendService) Update(ctx *gin.Context, cond *req.FriendUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.Friend{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.FriendRepo.Update(db, updateData)
}

func (s friendService) Delete(ctx *gin.Context, cond *req.FriendDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRepo.Delete(db, cond.ID)
}
