package service

import (
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/pkg/consts"
	"im/internal/util/ctxs"
	"im/internal/util/errs"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
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
		RequestStatusConds: []consts.FriendReqStatus{
			consts.FriendReqStatusPending,
			consts.FriendReqStatusRejected,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(fs.ID) != 0 {
		return nil, errs.RequestDuplicate
	}

	// 驗證對方是否發送過好友請求
	fs2, err := s.in.Repository.FriendRequestsRepo.Get(db, &req.FriendRequestsGet{
		SenderID:   u.ID,
		ReceiverID: loginID,
		RequestStatusConds: []consts.FriendReqStatus{
			consts.FriendReqStatusPending,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(fs2.ID) != 0 {
		// 雙方互相要求好友, 直接創建好友
		db = db.Begin()
		defer db.Rollback()
		err = s.createFriend(ctx, db, fs2)
		if err != nil {
			return nil, err
		}
		db.Commit()
		return fs2.ID, nil
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

	// 確認好友請求存在
	fr, err := s.in.Repository.FriendRequestsRepo.Get(db, &req.FriendRequestsGet{
		ID:         cond.ID,
		ReceiverID: ctxs.GetUserInfo(ctx).ID, // 必須是自己的ID
		RequestStatusConds: []consts.FriendReqStatus{
			consts.FriendReqStatusPending,
		},
	})
	if err != nil {
		return err
	}
	if len(fr.ID) == 0 {
		return errs.RequestInvalidID
	}

	db = db.Begin()
	defer db.Rollback()

	// 確認請求種類
	switch cond.RequestStatus {
	case consts.FriendReqStatusAccepted: // 接受
		err = s.createFriend(ctx, db, fr)
		if err != nil {
			return err
		}
	case consts.FriendReqStatusRejected: // 拒絕
		break
	default:
		return errs.CommonUnknownError
	}

	err = s.in.Repository.FriendRequestsRepo.Update(db, &models.FriendRequests{
		ID:            cond.ID,
		RequestStatus: consts.FriendReqStatus(cond.RequestStatus),
	})
	if err != nil {
		return err
	}
	db.Commit()
	return nil
}

func (s FriendRequestservice) Delete(ctx *gin.Context, cond *req.FriendRequestsDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.Delete(db, cond.ID)
}

// createFriend 創建好友
func (s FriendRequestservice) createFriend(ctx *gin.Context, db *gorm.DB, fr *models.FriendRequests) error {
	// 建立好友
	msgId := uuid.New()
	_, err := s.in.Repository.FriendRepo.Create(db, &models.Friend{
		ID:        uuid.New(),
		PUserID:   fr.SenderID,
		FUserID:   fr.ReceiverID,
		PUserName: fr.SenderName,
		FUserName: fr.ReceiverName,
		MessageId: msgId,
		Status:    consts.FriendStatusActive, Mute: false,
	})
	if err != nil {
		return err
	}
	_, err = s.in.Repository.FriendRepo.Create(db, &models.Friend{
		ID:        uuid.New(),
		PUserID:   fr.ReceiverID,
		FUserID:   fr.SenderID,
		PUserName: fr.ReceiverName,
		FUserName: fr.SenderName,
		MessageId: msgId,
		Status:    consts.FriendStatusActive, Mute: false,
	})
	if err != nil {
		return err
	}
	return nil
}
