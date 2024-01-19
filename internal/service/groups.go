package service

import (
	"im/internal/models/po"
	"im/internal/models/request"
	"im/internal/util/ctxs"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IGroupsService interface {
	Get(ctx *gin.Context, cond *request.GroupsGet) (*po.Groups, error)
	GetList(ctx *gin.Context, cond *request.GroupsGetList) (*po.PageResult[*po.Groups], error)
	Create(ctx *gin.Context, cond *request.GroupsCreate) (id any, err error)
	Update(ctx *gin.Context, cond *request.GroupsUpdate) (err error)
	Delete(ctx *gin.Context, cond *request.GroupsDelete) (err error)
}

func NewGroupsService(in DigIn) IGroupsService {
	return groupsService{in: in}
}

type groupsService struct {
	in DigIn
}

func (s groupsService) Get(ctx *gin.Context, cond *request.GroupsGet) (*po.Groups, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupsRepo.Get(db, cond)
}

func (s groupsService) GetList(ctx *gin.Context, cond *request.GroupsGetList) (*po.PageResult[*po.Groups], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupsRepo.GetList(db, cond)
}

func (s groupsService) Create(ctx *gin.Context, cond *request.GroupsCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &po.Groups{ID: uuid.New(), GroupOwnerID: ctxs.GetUserInfo(ctx).ID}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.GroupsRepo.Create(db, insertData)
}

func (s groupsService) Update(ctx *gin.Context, cond *request.GroupsUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &po.Groups{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.GroupsRepo.Update(db, updateData)
}

func (s groupsService) Delete(ctx *gin.Context, cond *request.GroupsDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupsRepo.Delete(db, cond.ID)
}
