package repository

import (
	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/pkg/sqldb"

	"github.com/coocood/freecache"
	"github.com/dtm-labs/rockscache"
	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
)

func NewRepository(in digIn) *Repository {
	return &Repository{in: in,
		ExampleRepo:         NewExampleRepository(in),
		UsersRepo:           NewUsersRepository(in),
		LoginRecordRepo:     NewLoginRecordRepository(in),
		FriendRepo:          NewFriendRepository(in),
		FriendRequestsRepo:  NewFriendRequestsRepository(in),
		GroupsRepo:          NewGroupsRepository(in),
		GroupMembersRepo:    NewGroupMembersRepository(in),
		CacheRepo:           NewCacheRepository(in),
		GroupInvitationRepo: NewGroupInvitationRepository(in),
		MessageRepo:         NewMessageRepository(in),
	}
}

type Repository struct {
	in digIn

	ExampleRepo         IExampleRepository
	UsersRepo           IUsersRepository
	LoginRecordRepo     ILoginRecordRepository
	FriendRepo          IFriendRepository
	FriendRequestsRepo  IFriendRequestsRepository
	GroupsRepo          IGroupsRepository
	GroupMembersRepo    IGroupMembersRepository
	CacheRepo           ICacheRepository
	GroupInvitationRepo IGroupInvitationRepository
	MessageRepo         IMessageRepository
}

// digIn repository require indendency
type digIn struct {
	dig.In

	Logger   logger.Logger
	Db       sqldb.Client
	Rdb      redis.UniversalClient
	Cache    *freecache.Cache
	Config   *config.Config
	RcClient *rockscache.Client
}
