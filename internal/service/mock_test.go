package service_test

import (
	"testing"

	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/pkg/sqldb"
	"im/internal/repository"
	"im/internal/repository/mock_repository"

	"github.com/stretchr/testify/mock"
	"go.uber.org/dig"
)

// MockRepoSet 包含所有仓库的模拟实例。
type MockRepoSet struct {
	ExampleRepo        *mock_repository.MockExampleRepository
	UsersRepo          *mock_repository.MockUsersRepository
	LoginRecordRepo    *mock_repository.MockLoginRecordRepository
	FriendRepo         *mock_repository.MockFriendRepository
	FriendRequestsRepo *mock_repository.MockFriendRequestsRepository
	GroupsRepo         *mock_repository.MockGroupsRepository
	GroupMembersRepo   *mock_repository.MockGroupMembersRepository
	RouteCacheRepo     *mock_repository.MockCacheRepository
}

// NewMockDigIn 创建一个具有所有模拟依赖项的 DigIn 对象。
func NewMockDigIn(t *testing.T) (*dig.Container, *MockRepoSet, *mock.Mock) {
	// 创建模拟仓库实例。
	mockRepo := &MockRepoSet{
		ExampleRepo:        new(mock_repository.MockExampleRepository),
		UsersRepo:          new(mock_repository.MockUsersRepository),
		LoginRecordRepo:    new(mock_repository.MockLoginRecordRepository),
		FriendRepo:         new(mock_repository.MockFriendRepository),
		FriendRequestsRepo: new(mock_repository.MockFriendRequestsRepository),
		GroupsRepo:         new(mock_repository.MockGroupsRepository),
		GroupMembersRepo:   new(mock_repository.MockGroupMembersRepository),
		RouteCacheRepo:     new(mock_repository.MockCacheRepository),
	}

	// 创建依赖注入容器。
	c := dig.New()

	// 提供所有模拟依赖。
	_ = c.Provide(func() *repository.Repository {
		return &repository.Repository{
			ExampleRepo:        mockRepo.ExampleRepo,
			UsersRepo:          mockRepo.UsersRepo,
			LoginRecordRepo:    mockRepo.LoginRecordRepo,
			FriendRepo:         mockRepo.FriendRepo,         // 添加模拟的Friend仓库
			FriendRequestsRepo: mockRepo.FriendRequestsRepo, // 添加模拟的FriendRequests仓库
			GroupsRepo:         mockRepo.GroupsRepo,         // 添加模拟的Groups仓库
			GroupMembersRepo:   mockRepo.GroupMembersRepo,   // 添加模拟的GroupMembers仓库
			CacheRepo:          mockRepo.RouteCacheRepo,     // 添加模拟的RouteCache仓库
		}
	})

	m := &mock.Mock{}

	// 提供日志记录器。
	_ = c.Provide(func() logger.Logger {
		return &logger.MockLogger{Mock: m}
	})

	// 提供数据库客户端。
	_ = c.Provide(sqldb.NewMockDB)

	// 提供配置。
	_ = c.Provide(func() *config.Config {
		// 返回你的测试配置或者模拟配置。
		return &config.Config{}
	})

	// 返回容器和模拟集合以供使用。
	return c, mockRepo, m
}
