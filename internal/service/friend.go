package service

import (
	"im/internal/models"
	"im/internal/models/request"
	"im/internal/pkg/consts/enums"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
)

type IFriendService interface {
	Get(ctx *gin.Context, cond *request.FriendGet) (*models.Friend, error)
	GetList(ctx *gin.Context, cond *request.FriendGetList) (*models.PageResult[*models.Friend], error)
	GetBlackList(ctx *gin.Context, cond *request.FriendGetList) (*models.PageResult[*models.Friend], error)
	GetMutualList(ctx *gin.Context, cond *request.FriendMutualGet) (*models.PageResult[*models.Friend], error)
	Update(ctx *gin.Context, cond *request.FriendUpdate) (err error)
	Delete(ctx *gin.Context, cond *request.FriendDelete) (err error)
}

func NewFriendService(in DigIn) IFriendService {
	return friendService{in: in}
}

type friendService struct {
	in DigIn
}

func (s friendService) Get(ctx *gin.Context, cond *request.FriendGet) (*models.Friend, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRepo.Get(db, cond)
}

func (s friendService) GetList(ctx *gin.Context, cond *request.FriendGetList) (*models.PageResult[*models.Friend], error) {
	db := s.in.DB.Session(ctx)
	cond.PUserID = ctxs.GetUserInfo(ctx).ID
	cond.Status = enums.FriendStatusActive
	return s.in.Repository.FriendRepo.GetList(db, cond)
}

func (s friendService) GetBlackList(ctx *gin.Context, cond *request.FriendGetList) (*models.PageResult[*models.Friend], error) {
	db := s.in.DB.Session(ctx)
	cond.PUserID = ctxs.GetUserInfo(ctx).ID
	cond.Status = enums.FriendStatusBlocked
	return s.in.Repository.FriendRepo.GetList(db, cond)
}

func (s friendService) GetMutualList(ctx *gin.Context, cond *request.FriendMutualGet) (*models.PageResult[*models.Friend], error) {
	db := s.in.DB.Session(ctx)
	cond.UserID = ctxs.GetUserInfo(ctx).ID
	return s.in.Repository.FriendRepo.GetMutualList(db, cond)
}

func (s friendService) Update(ctx *gin.Context, cond *request.FriendUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.Friend{
		PUserID: ctxs.GetUserInfo(ctx).ID,
		FUserID: cond.FUserID,
		Status:  cond.Status,
	}
	return s.in.Repository.FriendRepo.Update(db, updateData)
}

func (s friendService) Delete(ctx *gin.Context, cond *request.FriendDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRepo.Delete(db, cond.ID)
}
