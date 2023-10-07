package service

import (
	"im/internal/consts"
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/util/ctxs"
	"im/internal/util/errs"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IFriendRequestservice interface {
	Get(ctx *gin.Context, cond *req.FriendRequestsGet) (*models.FriendRequests, error)
	GetList(ctx *gin.Context, cond *req.FriendRequestsGetList) (*models.PageResult[*models.FriendRequests], error)
	Create(ctx *gin.Context, cond *req.FriendRequestsCreate) (id any, err error)
	Update(ctx *gin.Context, cond *req.FriendRequestsUpdate) (err error)
	Delete(ctx *gin.Context, cond *req.FriendRequestsDelete) (err error)
}

func NewFriendRequestservice(in digIn) IFriendRequestservice {
	return FriendRequestservice{in: in}
}

type FriendRequestservice struct {
	in digIn
}

func (s FriendRequestservice) Get(ctx *gin.Context, cond *req.FriendRequestsGet) (*models.FriendRequests, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.Get(db, cond)
}

func (s FriendRequestservice) GetList(ctx *gin.Context, cond *req.FriendRequestsGetList) (*models.PageResult[*models.FriendRequests], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.GetList(db, cond)
}

// Create 好友申請
func (s FriendRequestservice) Create(ctx *gin.Context, cond *req.FriendRequestsCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	loginID := ctxs.GetUserInfo(ctx).ID

	// 查詢對象ID
	u, err := s.in.Repository.UsersRepo.Get(db, &req.UsersGet{
		Username: cond.UserName,
	})
	if err != nil || len(u.ID) == 0 {
		return nil, errs.RequestInvalidUser
	}

	// 驗證是否已經是好友
	f, err := s.in.Repository.FriendRepo.Get(db, &req.FriendGet{
		PUserID: loginID,
		FUserID: u.ID,
	})
	if err != nil {
		return nil, err
	}
	if len(f.ID) != 0 {
		return nil, errs.BusinessFriendshipHint
	}

	// 驗證是否已經發出過好友請求
	fs, err := s.in.Repository.FriendRequestsRepo.Get(db, &req.FriendRequestsGet{
		SenderID:   loginID,
		ReceiverID: u.ID,
	})
	if err != nil {
		return nil, err
	}
	if len(fs.ID) != 0 {
		return nil, errs.RequestDuplicate
	}

	// 驗證對方是否發送過好友請求
	fs2, err := s.in.Repository.FriendRequestsRepo.Get(db, &req.FriendRequestsGet{
		SenderID:      u.ID,
		ReceiverID:    loginID,
		RequestStatus: consts.FriendReqStatusPending,
	})
	if err != nil {
		return nil, err
	}
	if len(fs2.ID) != 0 {
		// TODO 建立好友流程
	}

	// 建立好友請求
	insertData := &models.FriendRequests{
		ID:            uuid.New(),
		SenderID:      loginID,
		SenderName:    ctxs.GetUserInfo(ctx).Username,
		ReceiverID:    u.ID,
		ReceiverName:  u.Username,
		RequestStatus: consts.FriendReqStatusPending,
	}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.FriendRequestsRepo.Create(db, insertData)
}

func (s FriendRequestservice) Update(ctx *gin.Context, cond *req.FriendRequestsUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.FriendRequests{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.FriendRequestsRepo.Update(db, updateData)
}

func (s FriendRequestservice) Delete(ctx *gin.Context, cond *req.FriendRequestsDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.Delete(db, cond.ID)
}
