package user_test

import (
	"errors"
	"go-proj/internal/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the Repository interface
type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) GetByID(id string) (*user.User, error) {
	args := m.Called(id)
	if u := args.Get(0); u != nil {
		return u.(*user.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockRepo) Save(user *user.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Tests

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mockRepo)
	svc := user.NewService(mockRepo)

	mockRepo.
		On("Save", mock.AnythingOfType("*user.User")).
		Return(nil)

	userExpected, err := svc.CreateUser("Lucas", "lucas@example.com")

	assert.NoError(t, err)
	assert.NotNil(t, userExpected)
	assert.Equal(t, "Lucas", userExpected.Name)
	assert.Equal(t, "lucas@example.com", userExpected.Email)
	assert.NotEmpty(t, userExpected.Id)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_SaveFails(t *testing.T) {
	mockRepo := new(mockRepo)
	svc := user.NewService(mockRepo)

	expectedErr := errors.New("db error")

	mockRepo.
		On("Save", mock.AnythingOfType("*user.User")).
		Return(expectedErr)

	userExpected, err := svc.CreateUser("Lucas", "lucas@example.com")

	assert.Nil(t, userExpected)
	assert.EqualError(t, err, expectedErr.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetUser_Found(t *testing.T) {
	mockRepo := new(mockRepo)
	svc := user.NewService(mockRepo)

	expected := &user.User{
		Id:    "123",
		Name:  "Lucas",
		Email: "lucas@example.com",
	}

	mockRepo.
		On("GetByID", "123").
		Return(expected, nil)

	result, err := svc.GetUser("123")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	mockRepo := new(mockRepo)
	svc := user.NewService(mockRepo)

	mockRepo.
		On("GetByID", "notfound").
		Return(nil, errors.New("not found"))

	result, err := svc.GetUser("notfound")

	assert.Nil(t, result)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertExpectations(t)
}
