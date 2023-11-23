// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock_repository

import (
	models "im/internal/models"

	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	req "im/internal/models/req"
)

// MockFriendRequestsRepository is an autogenerated mock type for the IFriendRequestsRepository type
type MockFriendRequestsRepository struct {
	mock.Mock
}

type MockFriendRequestsRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFriendRequestsRepository) EXPECT() *MockFriendRequestsRepository_Expecter {
	return &MockFriendRequestsRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: db, data
func (_m *MockFriendRequestsRepository) Create(db *gorm.DB, data *models.FriendRequests) (interface{}, error) {
	ret := _m.Called(db, data)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.FriendRequests) (interface{}, error)); ok {
		return rf(db, data)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.FriendRequests) interface{}); ok {
		r0 = rf(db, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *models.FriendRequests) error); ok {
		r1 = rf(db, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFriendRequestsRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockFriendRequestsRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - db *gorm.DB
//   - data *models.FriendRequests
func (_e *MockFriendRequestsRepository_Expecter) Create(db interface{}, data interface{}) *MockFriendRequestsRepository_Create_Call {
	return &MockFriendRequestsRepository_Create_Call{Call: _e.mock.On("Create", db, data)}
}

func (_c *MockFriendRequestsRepository_Create_Call) Run(run func(db *gorm.DB, data *models.FriendRequests)) *MockFriendRequestsRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.FriendRequests))
	})
	return _c
}

func (_c *MockFriendRequestsRepository_Create_Call) Return(id interface{}, err error) *MockFriendRequestsRepository_Create_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockFriendRequestsRepository_Create_Call) RunAndReturn(run func(*gorm.DB, *models.FriendRequests) (interface{}, error)) *MockFriendRequestsRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: db, id
func (_m *MockFriendRequestsRepository) Delete(db *gorm.DB, id string) error {
	ret := _m.Called(db, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFriendRequestsRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockFriendRequestsRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - db *gorm.DB
//   - id string
func (_e *MockFriendRequestsRepository_Expecter) Delete(db interface{}, id interface{}) *MockFriendRequestsRepository_Delete_Call {
	return &MockFriendRequestsRepository_Delete_Call{Call: _e.mock.On("Delete", db, id)}
}

func (_c *MockFriendRequestsRepository_Delete_Call) Run(run func(db *gorm.DB, id string)) *MockFriendRequestsRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string))
	})
	return _c
}

func (_c *MockFriendRequestsRepository_Delete_Call) Return(err error) *MockFriendRequestsRepository_Delete_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFriendRequestsRepository_Delete_Call) RunAndReturn(run func(*gorm.DB, string) error) *MockFriendRequestsRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: db, cond
func (_m *MockFriendRequestsRepository) Get(db *gorm.DB, cond *req.FriendRequestsGet) (*models.FriendRequests, error) {
	ret := _m.Called(db, cond)

	var r0 *models.FriendRequests
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.FriendRequestsGet) (*models.FriendRequests, error)); ok {
		return rf(db, cond)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.FriendRequestsGet) *models.FriendRequests); ok {
		r0 = rf(db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.FriendRequests)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *req.FriendRequestsGet) error); ok {
		r1 = rf(db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFriendRequestsRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRouteCache'
type MockFriendRequestsRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - db *gorm.DB
//   - cond *req.FriendRequestsGet
func (_e *MockFriendRequestsRepository_Expecter) Get(db interface{}, cond interface{}) *MockFriendRequestsRepository_Get_Call {
	return &MockFriendRequestsRepository_Get_Call{Call: _e.mock.On("GetRouteCache", db, cond)}
}

func (_c *MockFriendRequestsRepository_Get_Call) Run(run func(db *gorm.DB, cond *req.FriendRequestsGet)) *MockFriendRequestsRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*req.FriendRequestsGet))
	})
	return _c
}

func (_c *MockFriendRequestsRepository_Get_Call) Return(_a0 *models.FriendRequests, _a1 error) *MockFriendRequestsRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFriendRequestsRepository_Get_Call) RunAndReturn(run func(*gorm.DB, *req.FriendRequestsGet) (*models.FriendRequests, error)) *MockFriendRequestsRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: db, cond
func (_m *MockFriendRequestsRepository) GetList(db *gorm.DB, cond *req.FriendRequestsGetList) (*models.PageResult[*models.FriendRequests], error) {
	ret := _m.Called(db, cond)

	var r0 *models.PageResult[*models.FriendRequests]
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.FriendRequestsGetList) (*models.PageResult[*models.FriendRequests], error)); ok {
		return rf(db, cond)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.FriendRequestsGetList) *models.PageResult[*models.FriendRequests]); ok {
		r0 = rf(db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PageResult[*models.FriendRequests])
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *req.FriendRequestsGetList) error); ok {
		r1 = rf(db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFriendRequestsRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockFriendRequestsRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - db *gorm.DB
//   - cond *req.FriendRequestsGetList
func (_e *MockFriendRequestsRepository_Expecter) GetList(db interface{}, cond interface{}) *MockFriendRequestsRepository_GetList_Call {
	return &MockFriendRequestsRepository_GetList_Call{Call: _e.mock.On("GetList", db, cond)}
}

func (_c *MockFriendRequestsRepository_GetList_Call) Run(run func(db *gorm.DB, cond *req.FriendRequestsGetList)) *MockFriendRequestsRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*req.FriendRequestsGetList))
	})
	return _c
}

func (_c *MockFriendRequestsRepository_GetList_Call) Return(_a0 *models.PageResult[*models.FriendRequests], _a1 error) *MockFriendRequestsRepository_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFriendRequestsRepository_GetList_Call) RunAndReturn(run func(*gorm.DB, *req.FriendRequestsGetList) (*models.PageResult[*models.FriendRequests], error)) *MockFriendRequestsRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: db, data
func (_m *MockFriendRequestsRepository) Update(db *gorm.DB, data *models.FriendRequests) error {
	ret := _m.Called(db, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.FriendRequests) error); ok {
		r0 = rf(db, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFriendRequestsRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockFriendRequestsRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - db *gorm.DB
//   - data *models.FriendRequests
func (_e *MockFriendRequestsRepository_Expecter) Update(db interface{}, data interface{}) *MockFriendRequestsRepository_Update_Call {
	return &MockFriendRequestsRepository_Update_Call{Call: _e.mock.On("Update", db, data)}
}

func (_c *MockFriendRequestsRepository_Update_Call) Run(run func(db *gorm.DB, data *models.FriendRequests)) *MockFriendRequestsRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.FriendRequests))
	})
	return _c
}

func (_c *MockFriendRequestsRepository_Update_Call) Return(err error) *MockFriendRequestsRepository_Update_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFriendRequestsRepository_Update_Call) RunAndReturn(run func(*gorm.DB, *models.FriendRequests) error) *MockFriendRequestsRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFriendRequestsRepository creates a new instance of MockFriendRequestsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFriendRequestsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFriendRequestsRepository {
	mock := &MockFriendRequestsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
