package http

import (
	"errors"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//userHandlerDynamoDb User handler dynamodb
type userHandlerDynamoDb struct {
	userService domain.UserServiceDynamoDb
}

//ProvideUserHandlerDynamoDb Provide user handler dynamodb
func ProvideUserHandlerDynamoDb(u domain.UserServiceDynamoDb) domain.UserHandlerDynamoDb {
	return userHandlerDynamoDb{userService: u}
}

// @BasePath /api/v1

// AddUser
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
// @Router /v1/users/ [post]
func (u userHandlerDynamoDb) Insert(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err = u.userService.Insert(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// @BasePath /api/v1

// AddUser
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
// @Router /v1/users/ [post]
func (u userHandlerDynamoDb) Find(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userService.Find(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (domain.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errors.New("user not found")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": user})
}

// @BasePath /api/v1

// Update User
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
// @Router /v1/users/ [put]
func (u userHandlerDynamoDb) Update(c *gin.Context) {
	var userDTO domain.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user, err := u.userService.Find(c, userDTO.UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (domain.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errors.New("user not found")})
		return
	}
	user.UUID = userDTO.UUID
	user.Username = userDTO.Username
	user.Password = userDTO.Password
	err = u.userService.Update(c, user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}

// @BasePath /api/v1

// DeleteUser
// @Summary Delete user
// @Schemes
// @Description Delete user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/{id} [delete]
func (u userHandlerDynamoDb) Delete(c *gin.Context) {
	id := c.Param("id")
	err := u.userService.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
