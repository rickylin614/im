package service

import (
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IGroupInvitationService interface {
	Get(ctx *gin.Context, cond *req.GroupInvitationGet) (*models.GroupInvitation, error)
	GetList(ctx *gin.Context, cond *req.GroupInvitationGetList) (*models.PageResult[*models.GroupInvitation], error)
	Create(ctx *gin.Context, cond *req.GroupInvitationCreate) (id any, err error)
	Update(ctx *gin.Context, cond *req.GroupInvitationUpdate) (err error)
	Delete(ctx *gin.Context, cond *req.GroupInvitationDelete) (err error)
}

func NewGroupInvitationService(in DigIn) IGroupInvitationService {
	return groupInvitationService{In: in}
}

type groupInvitationService struct {
	In DigIn
}

func (s groupInvitationService) Get(ctx *gin.Context, cond *req.GroupInvitationGet) (*models.GroupInvitation, error) {
	db := s.In.DB.Session(ctx)
	return s.In.Repository.GroupInvitationRepo.Get(db, cond)
}

func (s groupInvitationService) GetList(ctx *gin.Context, cond *req.GroupInvitationGetList) (*models.PageResult[*models.GroupInvitation], error) {
	db := s.In.DB.Session(ctx)
	return s.In.Repository.GroupInvitationRepo.GetList(db, cond)
}

func (s groupInvitationService) Create(ctx *gin.Context, cond *req.GroupInvitationCreate) (id any, err error) {
	db := s.In.DB.Session(ctx)
	insertData := &models.GroupInvitation{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.In.Repository.GroupInvitationRepo.Create(db, insertData)
}

func (s groupInvitationService) Update(ctx *gin.Context, cond *req.GroupInvitationUpdate) (err error) {
	db := s.In.DB.Session(ctx)
	updateData := &models.GroupInvitation{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.In.Repository.GroupInvitationRepo.Update(db, updateData)
}

func (s groupInvitationService) Delete(ctx *gin.Context, cond *req.GroupInvitationDelete) (err error) {
	db := s.In.DB.Session(ctx)
	return s.In.Repository.GroupInvitationRepo.Delete(db, cond.ID)
}
