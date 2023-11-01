package service

import (
	"go.uber.org/dig"

	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/pkg/sqldb"
	"im/internal/repository"
)

func NewService(in digIn) *Service {
	return &Service{in: in,
		ExampleSrv:       NewExampleService(in),
		UsersSrv:         NewUsersService(in),
		LoginRecordSrv:   NewLoginRecordService(in),
		FriendSrv:        NewFriendService(in),
		FriendRequestSrv: NewFriendRequestservice(in),
		GroupsSrv:        NewGroupsService(in),
		GroupMembersSrv:  NewGroupMembersService(in),
		RouteCacheSrv:    NewRouteCacheService(in),
	}
}

type Service struct {
	in digIn

	ExampleSrv       IExampleService
	UsersSrv         IUsersService
	LoginRecordSrv   ILoginRecordService
	FriendSrv        IFriendService
	FriendRequestSrv IFriendRequestservice
	GroupsSrv        IGroupsService
	GroupMembersSrv  IGroupMembersService
	RouteCacheSrv    IRouteCacheService
}

// digIn repository require independence
type digIn struct {
	dig.In

	Repository *repository.Repository
	Logger     logger.Logger
	DB         sqldb.Client
	Config     *config.Config
}
