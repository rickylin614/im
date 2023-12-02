package service

import (
	"time"

	"github.com/goccy/go-json"

	"im/internal/models"
	"im/internal/models/req"
	"im/internal/pkg/consts/rediskey"
	"im/internal/util/ctxs"
	"im/internal/util/errs"
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

func NewGroupMembersService(in DigIn) IGroupMembersService {
	return groupMembersService{in: in}
}

type groupMembersService struct {
	in DigIn
}

func (s groupMembersService) Get(ctx *gin.Context, cond *req.GroupMembersGet) (*models.GroupMembers, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.GroupMembersRepo.Get(db, cond)
}

func (s groupMembersService) GetList(ctx *gin.Context, cond *req.GroupMembersGetList) (*models.PageResult[*models.GroupMembers], error) {
	db := s.in.DB.Session(ctx)
	key := rediskey.GROUP_MEMBER_KEY + cond.Id

	result := &models.PageResult[*models.GroupMembers]{}
	// 取緩存
	buf, err := s.in.Repository.CacheRepo.GetCache(ctx, key)
	if err == nil {
		if err := json.Unmarshal(buf, result); err == nil {
			// 驗證是否為成員
			if s.IsGroupMember(ctx, result) {
				return result, nil
			} else {
				return nil, errs.RequestInvalidPermission
			}
		}
	}

	result, err = s.in.Repository.GroupMembersRepo.GetList(db, cond)

	// 設緩存
	if v, err := json.Marshal(result); err == nil {
		s.in.Repository.CacheRepo.SetCache(ctx, key, v, time.Hour)
	}

	// 驗證是否為成員
	if s.IsGroupMember(ctx, result) {
		return result, nil
	} else {
		return nil, errs.RequestInvalidPermission
	}
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

func (s groupMembersService) IsGroupMember(ctx *gin.Context, members *models.PageResult[*models.GroupMembers]) bool {
	for _, v := range members.Data {
		userid := ctxs.GetUserInfo(ctx).ID
		if v.UserID == userid {
			return true
		}
	}
	return false
}
