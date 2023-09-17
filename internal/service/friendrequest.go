package service

import (
	"context"
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/util/uuid"

	"github.com/jinzhu/copier"
)

type IFriendRequestservice interface {
	Get(ctx context.Context, cond *req.FriendRequestsGet) (*models.FriendRequests, error)
	GetList(ctx context.Context, cond *req.FriendRequestsGetList) (*models.PageResult[*models.FriendRequests], error)
	Create(ctx context.Context, cond *req.FriendRequestsCreate) (id any, err error)
	Update(ctx context.Context, cond *req.FriendRequestsUpdate) (err error)
	Delete(ctx context.Context, cond *req.FriendRequestsDelete) (err error)
}

func NewFriendRequestservice(in digIn) IFriendRequestservice {
	return FriendRequestservice{in: in}
}

type FriendRequestservice struct {
	in digIn
}

func (s FriendRequestservice) Get(ctx context.Context, cond *req.FriendRequestsGet) (*models.FriendRequests, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.Get(db, cond)
}

func (s FriendRequestservice) GetList(ctx context.Context, cond *req.FriendRequestsGetList) (*models.PageResult[*models.FriendRequests], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.GetList(db, cond)
}

func (s FriendRequestservice) Create(ctx context.Context, cond *req.FriendRequestsCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.FriendRequests{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.FriendRequestsRepo.Create(db, insertData)
}

func (s FriendRequestservice) Update(ctx context.Context, cond *req.FriendRequestsUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.FriendRequests{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.FriendRequestsRepo.Update(db, updateData)
}

func (s FriendRequestservice) Delete(ctx context.Context, cond *req.FriendRequestsDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.FriendRequestsRepo.Delete(db, cond.ID)
}
