package service_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"im/internal/models"
	"im/internal/models/req"
	"im/internal/pkg/consts/enums"
	"im/internal/pkg/logger"
	"im/internal/pkg/sqldb"
	"im/internal/repository/mock_repository"
	"im/internal/service"
	"im/internal/util/crypto"
	"im/internal/util/errs"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Mock dependencies
type mockDeps struct {
	UsersRepo *mock_repository.MockUsersRepository
	Logger    *logger.MockLogger
	DB        sqldb.Client
}

// Setup for your tests
func setupTest(t *testing.T) (service.IUsersService, *MockRepoSet, *mock.Mock) {
	container, mockRepoSet, mock := NewMockDigIn(t)

	var dIn service.DigIn
	err := container.Invoke(func(di service.DigIn) {
		dIn = di
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create the UsersService with the mock dependencies
	usersService := service.NewUsersService(dIn)

	// Return both the service and the mock dependencies
	return usersService, mockRepoSet, mock
}

type UserServiceTestSuite struct {
	suite.Suite
	UsersService service.IUsersService
	Deps         *MockRepoSet
	ctx          *gin.Context
	mock         *mock.Mock
}

// Before all tests
func (suite *UserServiceTestSuite) SetupSuite() {
	suite.UsersService, suite.Deps, suite.mock = setupTest(suite.T())

	gin.SetMode(gin.TestMode)
	request := httptest.NewRequest(http.MethodGet, "/users/login", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = request
	suite.ctx = ctx
}

// Before each test
func (suite *UserServiceTestSuite) SetupTest() {
	// setup for each test if needed
}

// After each test
func (suite *UserServiceTestSuite) TearDownTest() {
	// teardown for each test if needed
}

// After all tests
func (suite *UserServiceTestSuite) TearDownSuite() {
	suite.Deps.UsersRepo.AssertExpectations(suite.T())
	suite.Deps.LoginRecordRepo.AssertExpectations(suite.T())
}

func (suite *UserServiceTestSuite) TestGetUser() {
	ctx := &gin.Context{Request: &http.Request{}}
	getCond := &req.UsersGet{Username: "testuser"}
	expectedUser := &models.Users{Username: "testuser"}

	suite.Deps.UsersRepo.On("GetRouteCache", mock.Anything, getCond).Return(expectedUser, nil).Once()
	user, err := suite.UsersService.Get(ctx, getCond)

	suite.NoError(err)
	suite.Equal(expectedUser, user)
}

func (suite *UserServiceTestSuite) TestGetList() {
	ctx := &gin.Context{Request: &http.Request{}}
	getListCond := &req.UsersGetList{
		Page: models.Page{
			Index: 1,
			Size:  10,
		},
	}
	expectedResult := &models.PageResult[*models.Users]{
		Page: &models.Page{
			Index:     1,
			Size:      10,
			TotalPage: 1,
			Total:     2,
		},
		Data: []*models.Users{{Username: "testuser1"}, {Username: "testuser2"}},
	}

	suite.Deps.UsersRepo.On("GetList", mock.Anything, getListCond).Return(expectedResult, nil)
	result, err := suite.UsersService.GetList(ctx, getListCond)

	suite.NoError(err)
	suite.Equal(expectedResult, result)
}

func (suite *UserServiceTestSuite) TestCreate() {
	createCond := &req.UsersCreate{
		Username: "newuser",
		Password: "password123",
	}
	expectedId := uuid.New()

	suite.Deps.UsersRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Users")).Return(expectedId, nil).Once()
	id, err := suite.UsersService.Create(suite.ctx, createCond)

	suite.NoError(err)
	suite.Equal(expectedId, id)
}

func (suite *UserServiceTestSuite) TestUpdate() {
	updateCond := &req.UsersUpdate{
		ID:       uuid.New(),
		Nickname: "updatedNickname",
	}

	suite.Deps.UsersRepo.On("Update", mock.Anything, mock.AnythingOfType("*models.Users")).Return(nil)
	err := suite.UsersService.Update(suite.ctx, updateCond)

	suite.NoError(err)
}

func (suite *UserServiceTestSuite) TestDelete() {
	deleteCond := &req.UsersDelete{
		ID: uuid.New(),
	}

	suite.Deps.UsersRepo.On("Delete", mock.Anything, deleteCond.ID).Return(nil)
	err := suite.UsersService.Delete(suite.ctx, deleteCond)

	suite.NoError(err)
}

func (suite *UserServiceTestSuite) TestLogin() {
	loginCond := &req.UsersLogin{
		Username: "existinguser",
		Password: "password123",
	}
	expectedUser := &models.Users{Username: "existinguser", PasswordHash: crypto.Hash("password123")}

	suite.Deps.UsersRepo.On("GetRouteCache", mock.Anything, mock.AnythingOfType("*req.UsersGet")).Return(expectedUser, nil).Once()
	suite.Deps.UsersRepo.On("SetToken", suite.ctx, expectedUser.ID, mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
	suite.Deps.LoginRecordRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.LoginRecord")).Return(nil, nil).Once()
	token, err := suite.UsersService.Login(suite.ctx, loginCond)

	suite.NoError(err)
	suite.NotEmpty(token, "token should not empty")
}

func (suite *UserServiceTestSuite) TestLoginUserNotFound() {
	loginCond := &req.UsersLogin{
		Username: "nonexistinguser",
		Password: "password123",
	}

	suite.Deps.UsersRepo.On("GetRouteCache", mock.Anything, mock.AnythingOfType("*req.UsersGet")).Return(nil, errors.New("user not found")).Once()
	token, err := suite.UsersService.Login(suite.ctx, loginCond)

	suite.Error(err)
	suite.Empty(token)
	suite.Equal(errs.LoginCommonError, err)
}

func (suite *UserServiceTestSuite) TestLoginPasswordVerificationFailed() {
	loginCond := &req.UsersLogin{
		Username: "existinguser",
		Password: "wrongpassword",
	}
	expectedUser := &models.Users{Username: "existinguser", PasswordHash: crypto.Hash("correctpassword")}

	suite.Deps.UsersRepo.On("GetRouteCache", mock.Anything, mock.AnythingOfType("*req.UsersGet")).Return(expectedUser, nil).Once()
	suite.Deps.LoginRecordRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.LoginRecord")).Return(nil, nil).Once()
	token, err := suite.UsersService.Login(suite.ctx, loginCond)

	suite.Error(err)
	suite.Empty(token)
	suite.Equal(errs.LoginCommonError, err)
}

func (suite *UserServiceTestSuite) TestLoginUserStatusBlocked() {
	loginCond := &req.UsersLogin{
		Username: "blockeduser",
		Password: "password123",
	}
	expectedUser := &models.Users{Username: "blockeduser", PasswordHash: crypto.Hash("password123"), Status: enums.UserStatusBlocked}

	suite.Deps.UsersRepo.On("GetRouteCache", mock.Anything, mock.AnythingOfType("*req.UsersGet")).Return(expectedUser, nil).Once()
	suite.Deps.LoginRecordRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.LoginRecord")).Return(nil, nil).Once()
	token, err := suite.UsersService.Login(suite.ctx, loginCond)

	suite.Error(err)
	suite.Empty(token)
	suite.Equal(errs.LoginLockedError, err)
}

func (suite *UserServiceTestSuite) TestSetTokenFail() {
	loginCond := &req.UsersLogin{
		Username: "existinguser",
		Password: "password123",
	}
	expectedUser := &models.Users{Username: "existinguser", PasswordHash: crypto.Hash("password123")}

	suite.Deps.UsersRepo.On("GetRouteCache", mock.Anything, mock.AnythingOfType("*req.UsersGet")).Return(expectedUser, nil).Once()
	suite.Deps.UsersRepo.On("SetToken", suite.ctx, expectedUser.ID, mock.Anything, mock.AnythingOfType("string")).Return(errs.CommonServiceUnavailable).Once()
	suite.Deps.LoginRecordRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.LoginRecord")).Return(nil, nil).Once()
	suite.mock.On("Error", suite.ctx, fmt.Errorf("service set token err: %w", errs.CommonServiceUnavailable))
	token, err := suite.UsersService.Login(suite.ctx, loginCond)

	suite.Error(err)
	suite.Empty(token)
	suite.Equal(errs.CommonServiceUnavailable, err)
}

func (suite *UserServiceTestSuite) TestLogout() {
	ctx := &gin.Context{Request: &http.Request{}}
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXZpY2VfaWQiOiJEZWZhdWx0IiwidG9rZW4iOiIiLCJ1c2VyIjp7ImkiOiIiLCJ1IjoiZXhpc3Rpbmd1c2VyIiwibiI6IiIsInAiOiJiM2IwNTY4MGI3ZDZhMDBmZjIzMTU5YzBlZTE1NWQ5YyIsImUiOiIiLCJwaCI6IiIsImN0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJ1dCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicyI6MH0sImV4cCI6MTY5OTIwNDMzNCwianRpIjoiMDE4YmEwMGEtYTY3MS03MDNhLTliZDUtNThjMWRmNzQyNmJmIn0.SIIKD5LeijYEcn_tuu6Vpz33KpHqveSE_gHe_tza72hsgNXLfUekC9k8e1Qu3xGjR40_sSmiX6ePUTJ8CTtFHACE4Nz0TxW-Jx03WbhN5B3WyVL9lP3rkEyWlRwMgWPzSfk9nf0GbbDrjxnMX1obnkxjDzpG8QSC2WOR0p41JEA"

	suite.Deps.UsersRepo.On("DelToken", ctx, mock.Anything, mock.Anything).Return(nil)
	err := suite.UsersService.Logout(ctx, token)

	suite.NoError(err)
}

func (suite *UserServiceTestSuite) TestGetByToken() {
	ctx := &gin.Context{Request: &http.Request{}}
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXZpY2VfaWQiOiJEZWZhdWx0IiwidG9rZW4iOiIiLCJ1c2VyIjp7ImkiOiIiLCJ1IjoiZXhpc3Rpbmd1c2VyIiwibiI6IiIsInAiOiJiM2IwNTY4MGI3ZDZhMDBmZjIzMTU5YzBlZTE1NWQ5YyIsImUiOiIiLCJwaCI6IiIsImN0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJ1dCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwicyI6MH0sImV4cCI6MTY5OTIwNDMzNCwianRpIjoiMDE4YmEwMGEtYTY3MS03MDNhLTliZDUtNThjMWRmNzQyNmJmIn0.SIIKD5LeijYEcn_tuu6Vpz33KpHqveSE_gHe_tza72hsgNXLfUekC9k8e1Qu3xGjR40_sSmiX6ePUTJ8CTtFHACE4Nz0TxW-Jx03WbhN5B3WyVL9lP3rkEyWlRwMgWPzSfk9nf0GbbDrjxnMX1obnkxjDzpG8QSC2WOR0p41JEA"
	expectedJwtToken := &models.JWTClaims{}
	expectedJwtToken.User = &models.Users{ID: "", Username: "existinguser", PasswordHash: "b3b05680b7d6a00ff23159c0ee155d9c"}

	suite.Deps.UsersRepo.On("GetByToken", ctx, mock.Anything, mock.Anything, mock.Anything).Return(expectedJwtToken, nil)
	user, err := suite.UsersService.GetByToken(ctx, token)

	suite.NoError(err)
	suite.Equal(expectedJwtToken.User, user)
}

// This is the entry point for testing
func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
