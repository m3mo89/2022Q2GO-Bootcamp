package mock

import (
	"log"

	"github.com/stretchr/testify/mock"
)

type mockContext struct {
	mock.Mock
}

func NewContextMock() *mockContext {
	return &mockContext{}
}

func (context *mockContext) JSON(code int, i interface{}) error {
	log.Printf("Context Mock: JSON")
	arg := context.Called(code, i)
	return arg.Error(0)
}

func (context *mockContext) Bind(i interface{}) error {
	log.Printf("Context Mock: Bind")
	arg := context.Called(i)
	return arg.Error(0)
}

func (context *mockContext) Param(name string) string {
	log.Printf("Context Mock: Param")
	arg := context.Called(name)
	return arg.Get(0).(string)
}
