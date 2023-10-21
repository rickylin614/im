package service

import (
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IGroupsService interface {
	Get(ctx *gin.Context, cond *req.GroupsGet) (*models.Groups, error)
	GetList(ctx *gin.Context, cond *req.GroupsGetList) (*models.PageResult[*models.Groups], error)
	Create(ctx *gin.Context, cond *req.GroupsCreate) (id any, err error)
	Update(ctx *gin.Context, cond *req.GroupsUpdate) (err error)
	Delete(ctx *gin.Context, cond *req.GroupsDelete) (err error)
}

func NewGroupsService(in digIn) IGroupsService {
	return groupsService{in: in}
}

type groupsService struct {
	in digIn
}

func (s groupsService) Get(ctx *gin.Context, cond *req.GroupsGet) (*models.Groups, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupsRepo.Get(db, cond)
}

func (s groupsService) GetList(ctx *gin.Context, cond *req.GroupsGetList) (*models.PageResult[*models.Groups], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupsRepo.GetList(db, cond)
}

func (s groupsService) Create(ctx *gin.Context, cond *req.GroupsCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.Groups{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.GroupsRepo.Create(db, insertData)
}

func (s groupsService) Update(ctx *gin.Context, cond *req.GroupsUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.Groups{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.GroupsRepo.Update(db, updateData)
}

func (s groupsService) Delete(ctx *gin.Context, cond *req.GroupsDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupsRepo.Delete(db, cond.ID)
}
