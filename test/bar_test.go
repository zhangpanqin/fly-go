package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedApi struct {
	mock.Mock
}

func (m *MockedApi) get(name string) string {
	args := m.Called(name)
	return args.String(0)
}

func TestSomething(t *testing.T) {

	// create an instance of our test object
	testObj := new(MockedApi)

	// setup expectations
	testObj.On("get", mock.Anything).Return("567")

	// call the code we are testing
	get := Get("789", testObj)

	// assert that the expectations were met
	testObj.AssertExpectations(t)
	call := testObj.Calls[0]

	assert.Equal(t, "789567aaa", get)
	assert.Equal(t, "789", call.Arguments.String(0))
}
