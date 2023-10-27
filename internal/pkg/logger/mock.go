package logger

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Level() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockLogger) Debug(ctx context.Context, message string) {
	m.Called(ctx, message)
}

func (m *MockLogger) Info(ctx context.Context, message string) {
	m.Called(ctx, message)
}

func (m *MockLogger) Warn(ctx context.Context, message string) {
	m.Called(ctx, message)
}

func (m *MockLogger) Error(ctx context.Context, err error) {
	m.Called(ctx, err)
}

func (m *MockLogger) Panic(ctx context.Context, err error) {
	m.Called(ctx, err)
}

func (m *MockLogger) GetLogger() any {
	return m.Mock
}
