package service

import (
	"im/internal/models/po"
	"im/internal/models/request"
	"im/internal/util/ctxs"
	"im/internal/util/errs"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IGroupMembersService interface {
	Get(ctx *gin.Context, cond *request.GroupMembersGet) (*po.GroupMembers, error)
	GetList(ctx *gin.Context, cond *request.GroupMembersGetList) ([]*po.GroupMembers, error)
	Create(ctx *gin.Context, cond *request.GroupMembersCreate) (id any, err error)
	Update(ctx *gin.Context, cond *request.GroupMembersUpdate) (err error)
	Delete(ctx *gin.Context, cond *request.GroupMembersDelete) (err error)
}

func NewGroupMembersService(in DigIn) IGroupMembersService {
	return groupMembersService{in: in}
}

type groupMembersService struct {
	in DigIn
}

func (s groupMembersService) Get(ctx *gin.Context, cond *request.GroupMembersGet) (*po.GroupMembers, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupMembersRepo.Get(db, cond)
}

func (s groupMembersService) GetList(ctx *gin.Context, cond *request.GroupMembersGetList) ([]*po.GroupMembers, error) {
	db := s.in.DB.Session(ctx)

	result, err := s.in.Repository.GroupMembersRepo.GetListById(ctx, db, cond)
	if err != nil {
		return nil, err
	}

	// 驗證是否為成員
	if s.IsGroupMember(ctx, result) {
		return result, nil
	} else {
		return nil, errs.RequestInvalidPermission
	}
}

func (s groupMembersService) Create(ctx *gin.Context, cond *request.GroupMembersCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &po.GroupMembers{GroupID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.GroupMembersRepo.Create(db, insertData)
}

func (s groupMembersService) Update(ctx *gin.Context, cond *request.GroupMembersUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &po.GroupMembers{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.GroupMembersRepo.Update(db, updateData)
}

func (s groupMembersService) Delete(ctx *gin.Context, cond *request.GroupMembersDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupMembersRepo.Delete(db, cond.ID)
}

func (s groupMembersService) IsGroupMember(ctx *gin.Context, members []*po.GroupMembers) bool {
	for _, v := range members {
		userid := ctxs.GetUserInfo(ctx).ID
		if v.UserID == userid {
			return true
		}
	}
	return false
}
