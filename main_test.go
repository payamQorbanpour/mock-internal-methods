package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockOperation struct {
	mock.Mock
}

func (m *mockOperation) DoSomething(number int) bool {
	args := m.Called(number)
	return args.Bool(0)
}

func TestParrentFunction(t *testing.T) {
	mockOp := new(mockOperation)
	mockOp.On("DoSomething", 123).Return(true)

	testObj := &Operation{}
	done := testObj.ParrentFunction()

	assert.Equal(t, true, done)
}
