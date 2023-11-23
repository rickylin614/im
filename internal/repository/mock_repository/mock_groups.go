// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock_repository

import (
	models "im/internal/models"

	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	req "im/internal/models/req"
)

// MockGroupsRepository is an autogenerated mock type for the IGroupsRepository type
type MockGroupsRepository struct {
	mock.Mock
}

type MockGroupsRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGroupsRepository) EXPECT() *MockGroupsRepository_Expecter {
	return &MockGroupsRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: db, data
func (_m *MockGroupsRepository) Create(db *gorm.DB, data *models.Groups) (interface{}, error) {
	ret := _m.Called(db, data)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.Groups) (interface{}, error)); ok {
		return rf(db, data)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.Groups) interface{}); ok {
		r0 = rf(db, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *models.Groups) error); ok {
		r1 = rf(db, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupsRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockGroupsRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - db *gorm.DB
//   - data *models.Groups
func (_e *MockGroupsRepository_Expecter) Create(db interface{}, data interface{}) *MockGroupsRepository_Create_Call {
	return &MockGroupsRepository_Create_Call{Call: _e.mock.On("Create", db, data)}
}

func (_c *MockGroupsRepository_Create_Call) Run(run func(db *gorm.DB, data *models.Groups)) *MockGroupsRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.Groups))
	})
	return _c
}

func (_c *MockGroupsRepository_Create_Call) Return(id interface{}, err error) *MockGroupsRepository_Create_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockGroupsRepository_Create_Call) RunAndReturn(run func(*gorm.DB, *models.Groups) (interface{}, error)) *MockGroupsRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: db, id
func (_m *MockGroupsRepository) Delete(db *gorm.DB, id string) error {
	ret := _m.Called(db, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGroupsRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockGroupsRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - db *gorm.DB
//   - id string
func (_e *MockGroupsRepository_Expecter) Delete(db interface{}, id interface{}) *MockGroupsRepository_Delete_Call {
	return &MockGroupsRepository_Delete_Call{Call: _e.mock.On("Delete", db, id)}
}

func (_c *MockGroupsRepository_Delete_Call) Run(run func(db *gorm.DB, id string)) *MockGroupsRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string))
	})
	return _c
}

func (_c *MockGroupsRepository_Delete_Call) Return(err error) *MockGroupsRepository_Delete_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockGroupsRepository_Delete_Call) RunAndReturn(run func(*gorm.DB, string) error) *MockGroupsRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: db, cond
func (_m *MockGroupsRepository) Get(db *gorm.DB, cond *req.GroupsGet) (*models.Groups, error) {
	ret := _m.Called(db, cond)

	var r0 *models.Groups
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.GroupsGet) (*models.Groups, error)); ok {
		return rf(db, cond)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.GroupsGet) *models.Groups); ok {
		r0 = rf(db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Groups)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *req.GroupsGet) error); ok {
		r1 = rf(db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupsRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockGroupsRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - db *gorm.DB
//   - cond *req.GroupsGet
func (_e *MockGroupsRepository_Expecter) Get(db interface{}, cond interface{}) *MockGroupsRepository_Get_Call {
	return &MockGroupsRepository_Get_Call{Call: _e.mock.On("Get", db, cond)}
}

func (_c *MockGroupsRepository_Get_Call) Run(run func(db *gorm.DB, cond *req.GroupsGet)) *MockGroupsRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*req.GroupsGet))
	})
	return _c
}

func (_c *MockGroupsRepository_Get_Call) Return(_a0 *models.Groups, _a1 error) *MockGroupsRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupsRepository_Get_Call) RunAndReturn(run func(*gorm.DB, *req.GroupsGet) (*models.Groups, error)) *MockGroupsRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: db, cond
func (_m *MockGroupsRepository) GetList(db *gorm.DB, cond *req.GroupsGetList) (*models.PageResult[*models.Groups], error) {
	ret := _m.Called(db, cond)

	var r0 *models.PageResult[*models.Groups]
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.GroupsGetList) (*models.PageResult[*models.Groups], error)); ok {
		return rf(db, cond)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.GroupsGetList) *models.PageResult[*models.Groups]); ok {
		r0 = rf(db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PageResult[*models.Groups])
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *req.GroupsGetList) error); ok {
		r1 = rf(db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupsRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockGroupsRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - db *gorm.DB
//   - cond *req.GroupsGetList
func (_e *MockGroupsRepository_Expecter) GetList(db interface{}, cond interface{}) *MockGroupsRepository_GetList_Call {
	return &MockGroupsRepository_GetList_Call{Call: _e.mock.On("GetList", db, cond)}
}

func (_c *MockGroupsRepository_GetList_Call) Run(run func(db *gorm.DB, cond *req.GroupsGetList)) *MockGroupsRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*req.GroupsGetList))
	})
	return _c
}

func (_c *MockGroupsRepository_GetList_Call) Return(_a0 *models.PageResult[*models.Groups], _a1 error) *MockGroupsRepository_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupsRepository_GetList_Call) RunAndReturn(run func(*gorm.DB, *req.GroupsGetList) (*models.PageResult[*models.Groups], error)) *MockGroupsRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: db, data
func (_m *MockGroupsRepository) Update(db *gorm.DB, data *models.Groups) error {
	ret := _m.Called(db, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.Groups) error); ok {
		r0 = rf(db, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGroupsRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockGroupsRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - db *gorm.DB
//   - data *models.Groups
func (_e *MockGroupsRepository_Expecter) Update(db interface{}, data interface{}) *MockGroupsRepository_Update_Call {
	return &MockGroupsRepository_Update_Call{Call: _e.mock.On("Update", db, data)}
}

func (_c *MockGroupsRepository_Update_Call) Run(run func(db *gorm.DB, data *models.Groups)) *MockGroupsRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.Groups))
	})
	return _c
}

func (_c *MockGroupsRepository_Update_Call) Return(err error) *MockGroupsRepository_Update_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockGroupsRepository_Update_Call) RunAndReturn(run func(*gorm.DB, *models.Groups) error) *MockGroupsRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGroupsRepository creates a new instance of MockGroupsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGroupsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGroupsRepository {
	mock := &MockGroupsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
