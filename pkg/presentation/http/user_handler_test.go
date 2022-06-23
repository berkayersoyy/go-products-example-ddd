package http

//import (
//	"bytes"
//	"encoding/json"
//	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
//	"github.com/berkayersoyy/go-products-example-ddd/pkg/mocks"
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"github.com/twinj/uuid"
//	"log"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"time"
//)
//
//func TestUserHandler_GetAllUsers(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		w := httptest.NewRecorder()
//		c, _ := gin.CreateTestContext(w)
//		userUUID := uuid.NewV4()
//		users := []domain.User{{ID: "1", UUID: userUUID.String(), Username: "test-username", Password: "test-pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}}
//		mockApi := mocks.UserHandler{}
//		mockApi.On("GetAllUsers", mock.Anything).Return(gin.H{"users": users})
//
//		mockApi.GetAllUsers(c)
//		//c.Params = []gin.Param{{Key: "users", Value: "user1"}}
//
//		usersData := c.Param("users")
//		log.Println(usersData)
//		var got gin.H
//		err := json.Unmarshal(w.Body.Bytes(), &got)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 200, w.Code)
//		//assert.Equal(t, users, responseBody)
//		mockApi.AssertNumberOfCalls(t, "GetAllUsers", 1)
//	})
//
//}
//func TestUserHandler_GetUserByID(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		userUUID := uuid.NewV4()
//		users := []domain.User{{ID: "1", UUID: userUUID.String(), Username: "test-username", Password: "test-pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}}
//		mockApi := mocks.UserHandler{}
//		mockApi.On("GetUserByID", mock.Anything).Return(users)
//
//		w := httptest.NewRecorder()
//		r := gin.Default()
//		r.Use(func(context *gin.Context) {
//			context.Set("id", uint(1))
//		})
//		r.GET("/v1/users/:id", mockApi.GetUserByID)
//		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
//		req.Header.Set("Content-Type", "application/json")
//		assert.NoError(t, err)
//		r.ServeHTTP(w, req)
//
//		var respBody []domain.User
//		err = json.NewDecoder(w.Body).Decode(&respBody)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 200, w.Code)
//		assert.Equal(t, users, respBody)
//		mockApi.AssertNumberOfCalls(t, "GetUserByID", 1)
//	})
//
//}
//func TestUserHandler_AddUser(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		userUUID := uuid.NewV4()
//		user := domain.User{ID: "1", UUID: userUUID.String(), Username: "test-username", Password: "test-pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}
//		mockApi := mocks.UserHandler{}
//		mockApi.On("AddUser", mock.Anything).Return(user)
//
//		w := httptest.NewRecorder()
//		r := gin.Default()
//
//		reqBody, err := json.Marshal(user)
//		assert.NoError(t, err)
//		r.POST("/v1/users", mockApi.AddUser)
//		req, err := http.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		assert.NoError(t, err)
//		r.ServeHTTP(w, req)
//
//		var respBody domain.User
//		err = json.NewDecoder(w.Body).Decode(&respBody)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 200, w.Code)
//		assert.Equal(t, user, respBody)
//		mockApi.AssertNumberOfCalls(t, "AddUser", 1)
//	})
//}
//func TestUserHandler_UpdateUser(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		userUUID := uuid.NewV4()
//		user := domain.User{ID: "1", UUID: userUUID.String(), Username: "test-username", Password: "test-pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}
//		mockApi := mocks.UserHandler{}
//		mockApi.On("UpdateUser", mock.Anything).Return(user)
//
//		w := httptest.NewRecorder()
//		r := gin.Default()
//
//		reqBody, err := json.Marshal(user)
//		assert.NoError(t, err)
//		r.PUT("/v1/users", mockApi.UpdateUser)
//		req, err := http.NewRequest(http.MethodPut, "/v1/users", bytes.NewBuffer(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		assert.NoError(t, err)
//		r.ServeHTTP(w, req)
//
//		var respBody domain.User
//		err = json.NewDecoder(w.Body).Decode(&respBody)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 200, w.Code)
//		assert.Equal(t, user, respBody)
//		mockApi.AssertNumberOfCalls(t, "UpdateUser", 1)
//	})
//}
//func TestUserHandler_DeleteUser(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		mockApi := mocks.UserHandler{}
//		mockApi.On("DeleteUser", mock.Anything)
//
//		w := httptest.NewRecorder()
//		r := gin.Default()
//		r.Use(func(context *gin.Context) {
//			context.Set("id", uint(1))
//		})
//
//		r.DELETE("/v1/users/:id", mockApi.DeleteUser)
//		req, err := http.NewRequest(http.MethodDelete, "/v1/users/1", nil)
//		req.Header.Set("Content-Type", "application/json")
//		assert.NoError(t, err)
//		r.ServeHTTP(w, req)
//
//		assert.Equal(t, 200, w.Code)
//		mockApi.AssertNumberOfCalls(t, "DeleteUser", 1)
//	})
//}
