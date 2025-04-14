package user

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
	GetUserFn    func(id string) (*User, error)
	CreateUserFn func(name, email string) (*User, error)
}

func (m *mockService) GetUser(id string) (*User, error) {
	return m.GetUserFn(id)
}

func (m *mockService) CreateUser(name, email string) (*User, error) {
	return m.CreateUserFn(name, email)
}

func setupRouter(s UserService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	NewHandler(s).RegisterRoutes(r)
	return r
}

func TestGetUser_Success(t *testing.T) {
	mock := &mockService{
		GetUserFn: func(id string) (*User, error) {
			return &User{Id: id, Name: "user1", Email: "user1@test.com"}, nil
		},
	}

	router := setupRouter(mock)
	req, _ := http.NewRequest(http.MethodGet, "/users/123", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "user1@test.com")
}

func TestGetUser_NotFound(t *testing.T) {
	mock := &mockService{
		GetUserFn: func(id string) (*User, error) {
			return nil, errors.New("not found")
		},
	}

	router := setupRouter(mock)
	req, _ := http.NewRequest(http.MethodGet, "/users/999", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Contains(t, resp.Body.String(), "User not found")
}

func TestCreateUser_Success(t *testing.T) {
	mock := &mockService{
		CreateUserFn: func(name, email string) (*User, error) {
			return &User{Id: "1", Name: name, Email: email}, nil
		},
	}

	router := setupRouter(mock)
	body := []byte(`{"name":"user2","email":"user2@test.com"}`)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "user2@test.com")
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	mock := &mockService{}

	router := setupRouter(mock)
	body := []byte(`{invalid json}`)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid request")
}

func TestCreateUser_Failure(t *testing.T) {
	mock := &mockService{
		CreateUserFn: func(name, email string) (*User, error) {
			return nil, errors.New("db error")
		},
	}

	router := setupRouter(mock)
	body := []byte(`{"name":"user3","email":"user3@test.com"}`)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Contains(t, resp.Body.String(), "Could not create user")
}
