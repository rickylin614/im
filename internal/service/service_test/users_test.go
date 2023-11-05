package service_test

import (
	"net/http"
	"testing"

	"im/internal/models"
	"im/internal/models/req"
	"im/internal/pkg/logger"
	"im/internal/pkg/sqldb"
	"im/internal/repository"
	"im/internal/repository/mock_repository"
	"im/internal/service"

	// Adjust the import path according to your project

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock dependencies
type mockDeps struct {
	UsersRepo *mock_repository.MockUsersRepository
	Logger    *logger.MockLogger
	DB        *sqldb.MockClient
}

// Setup for your tests
func setupTest(t *testing.T) (*service.UsersService, *mockDeps) {
	// Initialize the mocks
	usersRepo := &mock_repository.MockUsersRepository{}
	logger := &logger.MockLogger{}
	db := &sqldb.MockClient{}

	// Create the digIn with the mocks
	in := service.In{
		Repository: &repository.Repository{UsersRepo: usersRepo},
		Logger:     logger,
		DB:         db,
	}

	// Create the UsersService with the mock dependencies
	usersService := service.NewUsersService(in)

	// Return both the service and the mock dependencies
	return usersService, &mockDeps{
		UsersRepo: usersRepo,
		Logger:    logger,
		DB:        db,
	}
}

// Example test for the Get method
func TestGetUser(t *testing.T) {
	usersService, deps := setupTest(t)
	ctx := &gin.Context{Request: &http.Request{}}

	// Mock inputs
	getCond := &req.UsersGet{Username: "testuser"}

	// Mock outputs
	expectedUser := &models.Users{Username: "testuser"}

	// Setup expectations
	deps.UsersRepo.On("Get", mock.Anything, getCond).Return(expectedUser, nil)

	// Call the method
	user, err := usersService.Get(ctx, getCond)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	// Assert that the expectations were met
	deps.UsersRepo.AssertExpectations(t)
}

// You can follow a similar pattern to test the rest of the methods.
