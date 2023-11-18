// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock_repository

import (
	context "context"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	models "im/internal/models"

	req "im/internal/models/req"
)

// MockUsersRepository is an autogenerated mock type for the IUsersRepository type
type MockUsersRepository struct {
	mock.Mock
}

type MockUsersRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUsersRepository) EXPECT() *MockUsersRepository_Expecter {
	return &MockUsersRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: db, data
func (_m *MockUsersRepository) Create(db *gorm.DB, data *models.Users) (interface{}, error) {
	ret := _m.Called(db, data)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.Users) (interface{}, error)); ok {
		return rf(db, data)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.Users) interface{}); ok {
		r0 = rf(db, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *models.Users) error); ok {
		r1 = rf(db, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUsersRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - db *gorm.DB
//   - data *models.Users
func (_e *MockUsersRepository_Expecter) Create(db interface{}, data interface{}) *MockUsersRepository_Create_Call {
	return &MockUsersRepository_Create_Call{Call: _e.mock.On("Create", db, data)}
}

func (_c *MockUsersRepository_Create_Call) Run(run func(db *gorm.DB, data *models.Users)) *MockUsersRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.Users))
	})
	return _c
}

func (_c *MockUsersRepository_Create_Call) Return(id interface{}, err error) *MockUsersRepository_Create_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockUsersRepository_Create_Call) RunAndReturn(run func(*gorm.DB, *models.Users) (interface{}, error)) *MockUsersRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DelToken provides a mock function with given fields: ctx, UserID, deviceID
func (_m *MockUsersRepository) DelToken(ctx context.Context, UserID string, deviceID string) error {
	ret := _m.Called(ctx, UserID, deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, UserID, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersRepository_DelToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DelToken'
type MockUsersRepository_DelToken_Call struct {
	*mock.Call
}

// DelToken is a helper method to define mock.On call
//   - ctx context.Context
//   - UserID string
//   - deviceID string
func (_e *MockUsersRepository_Expecter) DelToken(ctx interface{}, UserID interface{}, deviceID interface{}) *MockUsersRepository_DelToken_Call {
	return &MockUsersRepository_DelToken_Call{Call: _e.mock.On("DelToken", ctx, UserID, deviceID)}
}

func (_c *MockUsersRepository_DelToken_Call) Run(run func(ctx context.Context, UserID string, deviceID string)) *MockUsersRepository_DelToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockUsersRepository_DelToken_Call) Return(_a0 error) *MockUsersRepository_DelToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersRepository_DelToken_Call) RunAndReturn(run func(context.Context, string, string) error) *MockUsersRepository_DelToken_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: db, id
func (_m *MockUsersRepository) Delete(db *gorm.DB, id string) error {
	ret := _m.Called(db, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockUsersRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - db *gorm.DB
//   - id string
func (_e *MockUsersRepository_Expecter) Delete(db interface{}, id interface{}) *MockUsersRepository_Delete_Call {
	return &MockUsersRepository_Delete_Call{Call: _e.mock.On("Delete", db, id)}
}

func (_c *MockUsersRepository_Delete_Call) Run(run func(db *gorm.DB, id string)) *MockUsersRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string))
	})
	return _c
}

func (_c *MockUsersRepository_Delete_Call) Return(err error) *MockUsersRepository_Delete_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUsersRepository_Delete_Call) RunAndReturn(run func(*gorm.DB, string) error) *MockUsersRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: db, cond
func (_m *MockUsersRepository) Get(db *gorm.DB, cond *req.UsersGet) (*models.Users, error) {
	ret := _m.Called(db, cond)

	var r0 *models.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.UsersGet) (*models.Users, error)); ok {
		return rf(db, cond)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.UsersGet) *models.Users); ok {
		r0 = rf(db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Users)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *req.UsersGet) error); ok {
		r1 = rf(db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockUsersRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - db *gorm.DB
//   - cond *req.UsersGet
func (_e *MockUsersRepository_Expecter) Get(db interface{}, cond interface{}) *MockUsersRepository_Get_Call {
	return &MockUsersRepository_Get_Call{Call: _e.mock.On("Get", db, cond)}
}

func (_c *MockUsersRepository_Get_Call) Run(run func(db *gorm.DB, cond *req.UsersGet)) *MockUsersRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*req.UsersGet))
	})
	return _c
}

func (_c *MockUsersRepository_Get_Call) Return(_a0 *models.Users, _a1 error) *MockUsersRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersRepository_Get_Call) RunAndReturn(run func(*gorm.DB, *req.UsersGet) (*models.Users, error)) *MockUsersRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByToken provides a mock function with given fields: ctx, UserID, deviceID, reqToken
func (_m *MockUsersRepository) GetByToken(ctx context.Context, UserID string, deviceID string, reqToken string) (*models.JWTClaims, error) {
	ret := _m.Called(ctx, UserID, deviceID, reqToken)

	var r0 *models.JWTClaims
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*models.JWTClaims, error)); ok {
		return rf(ctx, UserID, deviceID, reqToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *models.JWTClaims); ok {
		r0 = rf(ctx, UserID, deviceID, reqToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JWTClaims)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, UserID, deviceID, reqToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersRepository_GetByToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByToken'
type MockUsersRepository_GetByToken_Call struct {
	*mock.Call
}

// GetByToken is a helper method to define mock.On call
//   - ctx context.Context
//   - UserID string
//   - deviceID string
//   - reqToken string
func (_e *MockUsersRepository_Expecter) GetByToken(ctx interface{}, UserID interface{}, deviceID interface{}, reqToken interface{}) *MockUsersRepository_GetByToken_Call {
	return &MockUsersRepository_GetByToken_Call{Call: _e.mock.On("GetByToken", ctx, UserID, deviceID, reqToken)}
}

func (_c *MockUsersRepository_GetByToken_Call) Run(run func(ctx context.Context, UserID string, deviceID string, reqToken string)) *MockUsersRepository_GetByToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockUsersRepository_GetByToken_Call) Return(_a0 *models.JWTClaims, _a1 error) *MockUsersRepository_GetByToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersRepository_GetByToken_Call) RunAndReturn(run func(context.Context, string, string, string) (*models.JWTClaims, error)) *MockUsersRepository_GetByToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: db, cond
func (_m *MockUsersRepository) GetList(db *gorm.DB, cond *req.UsersGetList) (*models.PageResult[*models.Users], error) {
	ret := _m.Called(db, cond)

	var r0 *models.PageResult[*models.Users]
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.UsersGetList) (*models.PageResult[*models.Users], error)); ok {
		return rf(db, cond)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.UsersGetList) *models.PageResult[*models.Users]); ok {
		r0 = rf(db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PageResult[*models.Users])
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *req.UsersGetList) error); ok {
		r1 = rf(db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockUsersRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - db *gorm.DB
//   - cond *req.UsersGetList
func (_e *MockUsersRepository_Expecter) GetList(db interface{}, cond interface{}) *MockUsersRepository_GetList_Call {
	return &MockUsersRepository_GetList_Call{Call: _e.mock.On("GetList", db, cond)}
}

func (_c *MockUsersRepository_GetList_Call) Run(run func(db *gorm.DB, cond *req.UsersGetList)) *MockUsersRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*req.UsersGetList))
	})
	return _c
}

func (_c *MockUsersRepository_GetList_Call) Return(_a0 *models.PageResult[*models.Users], _a1 error) *MockUsersRepository_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersRepository_GetList_Call) RunAndReturn(run func(*gorm.DB, *req.UsersGetList) (*models.PageResult[*models.Users], error)) *MockUsersRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// SetToken provides a mock function with given fields: ctx, UserID, deviceID, jwtData
func (_m *MockUsersRepository) SetToken(ctx context.Context, UserID string, deviceID string, jwtData string) error {
	ret := _m.Called(ctx, UserID, deviceID, jwtData)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, UserID, deviceID, jwtData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersRepository_SetToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetToken'
type MockUsersRepository_SetToken_Call struct {
	*mock.Call
}

// SetToken is a helper method to define mock.On call
//   - ctx context.Context
//   - UserID string
//   - deviceID string
//   - jwtData string
func (_e *MockUsersRepository_Expecter) SetToken(ctx interface{}, UserID interface{}, deviceID interface{}, jwtData interface{}) *MockUsersRepository_SetToken_Call {
	return &MockUsersRepository_SetToken_Call{Call: _e.mock.On("SetToken", ctx, UserID, deviceID, jwtData)}
}

func (_c *MockUsersRepository_SetToken_Call) Run(run func(ctx context.Context, UserID string, deviceID string, jwtData string)) *MockUsersRepository_SetToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockUsersRepository_SetToken_Call) Return(_a0 error) *MockUsersRepository_SetToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersRepository_SetToken_Call) RunAndReturn(run func(context.Context, string, string, string) error) *MockUsersRepository_SetToken_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: db, data
func (_m *MockUsersRepository) Update(db *gorm.DB, data *models.Users) error {
	ret := _m.Called(db, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.Users) error); ok {
		r0 = rf(db, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUsersRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - db *gorm.DB
//   - data *models.Users
func (_e *MockUsersRepository_Expecter) Update(db interface{}, data interface{}) *MockUsersRepository_Update_Call {
	return &MockUsersRepository_Update_Call{Call: _e.mock.On("Update", db, data)}
}

func (_c *MockUsersRepository_Update_Call) Run(run func(db *gorm.DB, data *models.Users)) *MockUsersRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.Users))
	})
	return _c
}

func (_c *MockUsersRepository_Update_Call) Return(err error) *MockUsersRepository_Update_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUsersRepository_Update_Call) RunAndReturn(run func(*gorm.DB, *models.Users) error) *MockUsersRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUsersRepository creates a new instance of MockUsersRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUsersRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUsersRepository {
	mock := &MockUsersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}