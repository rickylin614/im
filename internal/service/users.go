package service

import (
	"context"
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/util/crypto"
	"im/internal/util/uuid"

	"github.com/jinzhu/copier"
)

type IUsersService interface {
	Get(ctx context.Context, cond *req.UsersGet) (*models.Users, error)
	GetList(ctx context.Context, cond *req.UsersGetList) (*models.PageResult[*models.Users], error)
	Create(ctx context.Context, cond *req.UsersCreate) (id any, err error)
	Update(ctx context.Context, cond *req.UsersUpdate) (err error)
	Delete(ctx context.Context, cond *req.UsersDelete) (err error)
}

func NewUsersService(in digIn) IUsersService {
	return usersService{in: in}
}

type usersService struct {
	in digIn
}

func (s usersService) Get(ctx context.Context, cond *req.UsersGet) (*models.Users, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.UsersRepo.Get(db, cond)
}

func (s usersService) GetList(ctx context.Context, cond *req.UsersGetList) (*models.PageResult[*models.Users], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.UsersRepo.GetList(db, cond)
}

func (s usersService) Create(ctx context.Context, cond *req.UsersCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.Users{ID: uuid.New(), PasswordHash: crypto.Hash(cond.Password)}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.UsersRepo.Create(db, insertData)
}

func (s usersService) Update(ctx context.Context, cond *req.UsersUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.Users{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.UsersRepo.Update(db, updateData)
}

func (s usersService) Delete(ctx context.Context, cond *req.UsersDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.UsersRepo.Delete(db, cond.ID)
}
