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

//ProvideUserAPI Provide User Api
func ProvideUserAPI(u domain.UserService) domain.UserHandler {
	return &userHandler{UserService: u}
}

func (u *userHandler) GetAllUsers(c *gin.Context) {
	users := u.UserService.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *userHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.GetUserByID(uint(id))
	if user == (domain.User{}) {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": domain.ToUserDTO(user)})
}

func (u *userHandler) AddUser(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err)
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

func (u *userHandler) UpdateUser(c *gin.Context) {
	var userDTO domain.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		log.Println(err)
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
