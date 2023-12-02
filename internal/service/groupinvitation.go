package service

import (
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/pkg/consts/enums"
	"im/internal/util/ctxs"
	"im/internal/util/errs"
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
	// 檢查是群組成員 TODO 創一層共用代碼
	m, err := s.In.Repository.GroupMembersRepo.Get(db, &req.GroupMembersGet{
		UserId: ctxs.GetUserInfo(ctx).ID,
	})
	if err != nil {
		return nil, err
	} else if m == nil {
		return nil, errs.RequestInvalidPermission
	}

	// 對象已是群組成員
	t, err := s.In.Repository.GroupMembersRepo.Get(db, &req.GroupMembersGet{
		UserId: cond.InviteeId,
	})
	if err != nil {
		return nil, err
	} else if t != nil {
		return nil, errs.GroupMemberExistError
	}

	// 檢查是否邀請過
	i, err := s.In.Repository.GroupInvitationRepo.Get(db, &req.GroupInvitationGet{
		GroupID:   cond.GroupId,
		InviterID: ctxs.GetUserInfo(ctx).ID,
		InviteeID: cond.InviteeId,
	})
	if err != nil {
		return nil, err
	} else if i != nil {
		return nil, errs.RequestDuplicate
	}

	// 創建邀請
	insertData := &models.GroupInvitation{
		ID:               uuid.New(),
		GroupID:          cond.GroupId,
		InviterID:        ctxs.GetUserInfo(ctx).ID,
		InviteeID:        cond.InviteeId,
		InvitationStatus: enums.GroupInvitationStatusPending,
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
