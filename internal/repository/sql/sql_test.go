package sql

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"im/internal/pkg/logger"
)

type SQLSuite struct {
	suite.Suite
	Files      *Files
	MockFS     *MockFS
	MockLogger *logger.MockLogger
}

func (s *SQLSuite) SetupTest() {
	// 初始化模拟对象
	s.MockFS = new(MockFS)
	s.MockLogger = &logger.MockLogger{Mock: &mock.Mock{}}

	// 初始化Files
	s.Files = &Files{
		in:         digIn{Logger: s.MockLogger},
		embedFiles: s.MockFS,
	}
}

func (s *SQLSuite) TestLoadSuccessfulRead() {
	s.MockFS.On("ReadFile", "test.sql").Return([]byte("SELECT * FROM test;"), nil)
	s.Equal("SELECT * FROM test;", s.Files.Load("test.sql"))
}

func (s *SQLSuite) TestLoadFileNotFound() {
	s.MockFS.On("ReadFile", "notfound.sql").Return([]byte{}, errors.New("file not found"))
	s.MockLogger.On("Error", mock.Anything, mock.Anything).Once()
	s.Equal("", s.Files.Load("notfound.sql"))
	s.MockLogger.AssertExpectations(s.T())
}

func TestSQLSuite(t *testing.T) {
	suite.Run(t, new(SQLSuite))
}

type MockFS struct {
	mock.Mock
}

func (m *MockFS) ReadFile(name string) ([]byte, error) {
	args := m.Called(name)
	return args.Get(0).([]byte), args.Error(1)
}
