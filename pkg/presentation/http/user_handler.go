package http

import (
	"context"
	"errors"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/tracing/jaeger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
	"net/http"
)

//userHandler User handler dynamodb
type userHandler struct {
	userService domain.UserService
}

//ProvideUserHandler Provide user handler dynamodb
func ProvideUserHandler(u domain.UserService) domain.UserHandler {
	return userHandler{userService: u}
}

// @BasePath /api/v1

// Insert
// @Summary Add user
// @Schemes
// @Description Add user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.User true "User ID"
// @Success 200 {object} domain.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/dynamodb/users/ [post]
func (u userHandler) Insert(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "UserHandler.Insert", "POST")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	defer span.Finish()
	ctx := context.WithValue(c, "header", c.Request.Header)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		log.Fatalf("Error %s", err)
	}
	var user domain.User
	err = c.BindJSON(&user)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err = u.userService.Insert(ctx, user)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	span.SetTag("http.status_code", "201")
	c.Status(http.StatusCreated)
}

// @BasePath /api/v1

// FindByUUID
// @Summary Find user
// @Schemes
// @Description Find user by uuid
// @Tags Users
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {object} domain.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/dynamodb/users/getbyuuid/{uuid} [get]
func (u userHandler) FindByUUID(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "UserHandler.FindByUUID", "GET")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	defer span.Finish()
	ctx := context.WithValue(c, "header", c.Request.Header)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		log.Fatalf("Error %s", err)
	}
	id := c.Param("uuid")
	user, err := u.userService.FindByUUID(ctx, id)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (domain.User{}) {
		errStr := errors.New("user not found")
		ext.LogError(span, errStr)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": errStr})
		return
	}
	span.SetTag("http.status_code", "200")
	c.JSON(http.StatusOK, gin.H{"User": user})
}

// @BasePath /api/v1

// FindByUsername
// @Summary Find user
// @Schemes
// @Description Find user by username
// @Tags Users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} domain.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/dynamodb/users/getbyusername/{username} [get]
func (u userHandler) FindByUsername(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "UserHandler.FindByUsername", "GET")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	defer span.Finish()
	ctx := context.WithValue(c, "header", c.Request.Header)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		log.Fatalf("Error %s", err)
	}
	username := c.Param("username")
	user, err := u.userService.FindByUsername(ctx, username)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (domain.User{}) {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": errors.New("user not found")})
		return
	}
	span.SetTag("http.status_code", "200")
	c.JSON(http.StatusOK, gin.H{"User": user})
}

// @BasePath /api/v1

// Update
// @Summary Update user
// @Schemes
// @Description Update user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.UserDTO true "User ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/dynamodb/users/ [put]
func (u userHandler) Update(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "UserHandler.Update", "PUT")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	defer span.Finish()
	ctx := context.WithValue(c, "header", c.Request.Header)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		log.Fatalf("Error %s", err)
	}
	var userDTO domain.UserDTO
	err = c.BindJSON(&userDTO)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(userDTO)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user, err := u.userService.FindByUUID(ctx, userDTO.UUID)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (domain.User{}) {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": errors.New("user not found")})
		return
	}
	user.UUID = userDTO.UUID
	user.Username = userDTO.Username
	user.Password = userDTO.Password
	err = u.userService.Update(ctx, user)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.Status(http.StatusBadRequest)
		return
	}
	span.SetTag("http.status_code", "201")
	c.Status(http.StatusCreated)
}

// @BasePath /api/v1

// Delete
// @Summary Delete user
// @Schemes
// @Description Delete user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User UUID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/dynamodb/users/{id} [delete]
func (u userHandler) Delete(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "UserHandler.Delete", "DELETE")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	defer span.Finish()
	ctx := context.WithValue(c, "header", c.Request.Header)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		log.Fatalf("Error %s", err)
	}
	id := c.Param("uuid")
	err = u.userService.Delete(ctx, id)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	span.SetTag("http.status_code", "201")
	c.Status(http.StatusCreated)
}
