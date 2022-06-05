package http

import (
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

type userHandler struct {
	UserService domain.UserService
}

func ProvideuserAPI(u domain.UserService) domain.UserHandler {
	return &userHandler{UserService: u}
}

// @BasePath /api/v1

// GetAllUsers
// @Summary Fetch all users
// @Schemes
// @Description Fetch all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} domain.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/ [get]
func (u *userHandler) GetAllUsers(c *gin.Context) {
	users := u.UserService.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// @BasePath /api/v1

// GetUserByID
// @Summary Fetch user by id
// @Schemes
// @Description Fetch user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/{id} [get]
func (u *userHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.GetUserByID(uint(id))
	if user == (domain.User{}) {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": domain.ToUserDTO(user)})
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
func (u *userHandler) AddUser(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	createdUser := u.UserService.AddUser(user)
	c.JSON(http.StatusOK, gin.H{"user": domain.ToUserDTO(createdUser)})
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
func (u *userHandler) UpdateUser(c *gin.Context) {
	var userDTO domain.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.GetUserByID(uint(id))
	if user == (domain.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	user.Username = userDTO.Username
	user.Password = userDTO.Password
	u.UserService.AddUser(user)

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
func (u *userHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.GetUserByID(uint(id))
	if user == (domain.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	u.UserService.DeleteUser(user)
	c.Status(http.StatusOK)
}
func (u *userHandler) GetUserByUsername(c *gin.Context) {

	un := c.Param("username")
	user := u.UserService.GetUserByUsername(un)
	if (user == domain.User{}) {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": domain.ToUserDTO(user)})

}
