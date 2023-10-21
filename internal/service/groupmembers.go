package service

import (
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IGroupMembersService interface {
	Get(ctx *gin.Context, cond *req.GroupMembersGet) (*models.GroupMembers, error)
	GetList(ctx *gin.Context, cond *req.GroupMembersGetList) (*models.PageResult[*models.GroupMembers], error)
	Create(ctx *gin.Context, cond *req.GroupMembersCreate) (id any, err error)
	Update(ctx *gin.Context, cond *req.GroupMembersUpdate) (err error)
	Delete(ctx *gin.Context, cond *req.GroupMembersDelete) (err error)
}

func NewGroupMembersService(in digIn) IGroupMembersService {
	return groupMembersService{in: in}
}

type groupMembersService struct {
	in digIn
}

func (s groupMembersService) Get(ctx *gin.Context, cond *req.GroupMembersGet) (*models.GroupMembers, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupMembersRepo.Get(db, cond)
}

func (s groupMembersService) GetList(ctx *gin.Context, cond *req.GroupMembersGetList) (*models.PageResult[*models.GroupMembers], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupMembersRepo.GetList(db, cond)
}

func (s groupMembersService) Create(ctx *gin.Context, cond *req.GroupMembersCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.GroupMembers{GroupID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.GroupMembersRepo.Create(db, insertData)
}

func (s groupMembersService) Update(ctx *gin.Context, cond *req.GroupMembersUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.GroupMembers{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.GroupMembersRepo.Update(db, updateData)
}

func (s groupMembersService) Delete(ctx *gin.Context, cond *req.GroupMembersDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupMembersRepo.Delete(db, cond.ID)
}
