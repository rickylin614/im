package service

import (
	"im/internal/models/po"
	"im/internal/models/request"
	"im/internal/pkg/consts/enums"
	"im/internal/util/ctxs"
	"im/internal/util/errs"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type IFriendRequestservice interface {
	Get(ctx *gin.Context, cond *request.FriendRequestsGet) (*po.FriendRequests, error)
	GetList(ctx *gin.Context, cond *request.FriendRequestsGetList) (*po.PageResult[*po.FriendRequests], error)
	Create(ctx *gin.Context, cond *request.FriendRequestsCreate) (id any, err error)
	Update(ctx *gin.Context, cond *request.FriendRequestsUpdate) (err error)
	Delete(ctx *gin.Context, cond *request.FriendRequestsDelete) (err error)
}

func NewFriendRequestservice(in DigIn) IFriendRequestservice {
	return FriendRequestservice{in: in}
}

type FriendRequestservice struct {
	in DigIn
}

func (s FriendRequestservice) Get(ctx *gin.Context, cond *request.FriendRequestsGet) (*po.FriendRequests, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.Get(db, cond)
}

func (s FriendRequestservice) GetList(ctx *gin.Context, cond *request.FriendRequestsGetList) (*po.PageResult[*po.FriendRequests], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.GetList(db, cond)
}

// Create 好友申請
func (s FriendRequestservice) Create(ctx *gin.Context, cond *request.FriendRequestsCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	loginID := ctxs.GetUserInfo(ctx).ID

	// 查詢對象ID
	u, err := s.in.Repository.UsersRepo.Get(db, &request.UsersGet{
		Username: cond.UserName,
	})
	if err != nil || len(u.ID) == 0 {
		return nil, errs.RequestInvalidUser
	}

	// 驗證是否已經是好友
	f, err := s.in.Repository.FriendRepo.Get(db, &request.FriendGet{
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
	fs, err := s.in.Repository.FriendRequestsRepo.Get(db, &request.FriendRequestsGet{
		SenderID:   loginID,
		ReceiverID: u.ID,
		RequestStatusConds: []enums.FriendReqStatus{
			enums.FriendReqStatusPending,
			enums.FriendReqStatusRejected,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(fs.ID) != 0 {
		return nil, errs.RequestDuplicate
	}

	// 驗證對方是否發送過好友請求
	fs2, err := s.in.Repository.FriendRequestsRepo.Get(db, &request.FriendRequestsGet{
		SenderID:   u.ID,
		ReceiverID: loginID,
		RequestStatusConds: []enums.FriendReqStatus{
			enums.FriendReqStatusPending,
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
	insertData := &po.FriendRequests{
		ID:            uuid.New(),
		SenderID:      loginID,
		SenderName:    ctxs.GetUserInfo(ctx).Username,
		ReceiverID:    u.ID,
		ReceiverName:  u.Username,
		RequestStatus: enums.FriendReqStatusPending,
	}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.FriendRequestsRepo.Create(db, insertData)
}

func (s FriendRequestservice) Update(ctx *gin.Context, cond *request.FriendRequestsUpdate) (err error) {
	db := s.in.DB.Session(ctx)

	// 確認好友請求存在
	fr, err := s.in.Repository.FriendRequestsRepo.Get(db, &request.FriendRequestsGet{
		ID:         cond.ID,
		ReceiverID: ctxs.GetUserInfo(ctx).ID, // 必須是自己的ID
		RequestStatusConds: []enums.FriendReqStatus{
			enums.FriendReqStatusPending,
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
	case enums.FriendReqStatusAccepted: // 接受
		err = s.createFriend(ctx, db, fr)
		if err != nil {
			return err
		}
	case enums.FriendReqStatusRejected: // 拒絕
		break
	default:
		return errs.CommonUnknownError
	}

	err = s.in.Repository.FriendRequestsRepo.Update(db, &po.FriendRequests{
		ID:            cond.ID,
		RequestStatus: enums.FriendReqStatus(cond.RequestStatus),
	})
	if err != nil {
		return err
	}
	db.Commit()
	return nil
}

func (s FriendRequestservice) Delete(ctx *gin.Context, cond *request.FriendRequestsDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.Delete(db, cond.ID)
}

// createFriend 創建好友
func (s FriendRequestservice) createFriend(ctx *gin.Context, db *gorm.DB, fr *po.FriendRequests) error {
	// 建立好友
	msgId := uuid.New()
	_, err := s.in.Repository.FriendRepo.Create(db, &po.Friend{
		ID:        uuid.New(),
		PUserID:   fr.SenderID,
		FUserID:   fr.ReceiverID,
		PUserName: fr.SenderName,
		FUserName: fr.ReceiverName,
		MessageId: msgId,
		Status:    enums.FriendStatusActive, Mute: false,
	})
	if err != nil {
		return err
	}
	_, err = s.in.Repository.FriendRepo.Create(db, &po.Friend{
		ID:        uuid.New(),
		PUserID:   fr.ReceiverID,
		FUserID:   fr.SenderID,
		PUserName: fr.ReceiverName,
		FUserName: fr.SenderName,
		MessageId: msgId,
		Status:    enums.FriendStatusActive, Mute: false,
	})
	if err != nil {
		return err
	}
	return nil
}
