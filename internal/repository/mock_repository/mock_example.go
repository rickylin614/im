// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock_repository

import (
	"im/internal/models/po"

	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	req "im/internal/models/request"
)

// MockExampleRepository is an autogenerated mock type for the IExampleRepository type
type MockExampleRepository struct {
	mock.Mock
}

type MockExampleRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExampleRepository) EXPECT() *MockExampleRepository_Expecter {
	return &MockExampleRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: db, data
func (_m *MockExampleRepository) Create(db *gorm.DB, data *po.Example) (interface{}, error) {
	ret := _m.Called(db, data)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *po.Example) (interface{}, error)); ok {
		return rf(db, data)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *po.Example) interface{}); ok {
		r0 = rf(db, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *po.Example) error); ok {
		r1 = rf(db, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExampleRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockExampleRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - db *gorm.DB
//   - data *models.Example
func (_e *MockExampleRepository_Expecter) Create(db interface{}, data interface{}) *MockExampleRepository_Create_Call {
	return &MockExampleRepository_Create_Call{Call: _e.mock.On("Create", db, data)}
}

func (_c *MockExampleRepository_Create_Call) Run(run func(db *gorm.DB, data *po.Example)) *MockExampleRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*po.Example))
	})
	return _c
}

func (_c *MockExampleRepository_Create_Call) Return(id interface{}, err error) *MockExampleRepository_Create_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockExampleRepository_Create_Call) RunAndReturn(run func(*gorm.DB, *po.Example) (interface{}, error)) *MockExampleRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: db, id
func (_m *MockExampleRepository) Delete(db *gorm.DB, id string) error {
	ret := _m.Called(db, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExampleRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockExampleRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - db *gorm.DB
//   - id string
func (_e *MockExampleRepository_Expecter) Delete(db interface{}, id interface{}) *MockExampleRepository_Delete_Call {
	return &MockExampleRepository_Delete_Call{Call: _e.mock.On("Delete", db, id)}
}

func (_c *MockExampleRepository_Delete_Call) Run(run func(db *gorm.DB, id string)) *MockExampleRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string))
	})
	return _c
}

func (_c *MockExampleRepository_Delete_Call) Return(err error) *MockExampleRepository_Delete_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockExampleRepository_Delete_Call) RunAndReturn(run func(*gorm.DB, string) error) *MockExampleRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: db, cond
func (_m *MockExampleRepository) Get(db *gorm.DB, cond *req.ExampleGet) (*po.Example, error) {
	ret := _m.Called(db, cond)

	var r0 *po.Example
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.ExampleGet) (*po.Example, error)); ok {
		return rf(db, cond)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.ExampleGet) *po.Example); ok {
		r0 = rf(db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*po.Example)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *req.ExampleGet) error); ok {
		r1 = rf(db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExampleRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockExampleRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - db *gorm.DB
//   - cond *request.ExampleGet
func (_e *MockExampleRepository_Expecter) Get(db interface{}, cond interface{}) *MockExampleRepository_Get_Call {
	return &MockExampleRepository_Get_Call{Call: _e.mock.On("Get", db, cond)}
}

func (_c *MockExampleRepository_Get_Call) Run(run func(db *gorm.DB, cond *req.ExampleGet)) *MockExampleRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*req.ExampleGet))
	})
	return _c
}

func (_c *MockExampleRepository_Get_Call) Return(_a0 *po.Example, _a1 error) *MockExampleRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExampleRepository_Get_Call) RunAndReturn(run func(*gorm.DB, *req.ExampleGet) (*po.Example, error)) *MockExampleRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: db, cond
func (_m *MockExampleRepository) GetList(db *gorm.DB, cond *req.ExampleGetList) (*po.PageResult[*po.Example], error) {
	ret := _m.Called(db, cond)

	var r0 *po.PageResult[*po.Example]
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.ExampleGetList) (*po.PageResult[*po.Example], error)); ok {
		return rf(db, cond)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *req.ExampleGetList) *po.PageResult[*po.Example]); ok {
		r0 = rf(db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*po.PageResult[*po.Example])
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *req.ExampleGetList) error); ok {
		r1 = rf(db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExampleRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockExampleRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - db *gorm.DB
//   - cond *request.ExampleGetList
func (_e *MockExampleRepository_Expecter) GetList(db interface{}, cond interface{}) *MockExampleRepository_GetList_Call {
	return &MockExampleRepository_GetList_Call{Call: _e.mock.On("GetList", db, cond)}
}

func (_c *MockExampleRepository_GetList_Call) Run(run func(db *gorm.DB, cond *req.ExampleGetList)) *MockExampleRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*req.ExampleGetList))
	})
	return _c
}

func (_c *MockExampleRepository_GetList_Call) Return(_a0 *po.PageResult[*po.Example], _a1 error) *MockExampleRepository_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExampleRepository_GetList_Call) RunAndReturn(run func(*gorm.DB, *req.ExampleGetList) (*po.PageResult[*po.Example], error)) *MockExampleRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: db, data
func (_m *MockExampleRepository) Update(db *gorm.DB, data *po.Example) error {
	ret := _m.Called(db, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *po.Example) error); ok {
		r0 = rf(db, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExampleRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockExampleRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - db *gorm.DB
//   - data *models.Example
func (_e *MockExampleRepository_Expecter) Update(db interface{}, data interface{}) *MockExampleRepository_Update_Call {
	return &MockExampleRepository_Update_Call{Call: _e.mock.On("Update", db, data)}
}

func (_c *MockExampleRepository_Update_Call) Run(run func(db *gorm.DB, data *po.Example)) *MockExampleRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*po.Example))
	})
	return _c
}

func (_c *MockExampleRepository_Update_Call) Return(err error) *MockExampleRepository_Update_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockExampleRepository_Update_Call) RunAndReturn(run func(*gorm.DB, *po.Example) error) *MockExampleRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExampleRepository creates a new instance of MockExampleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExampleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExampleRepository {
	mock := &MockExampleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
