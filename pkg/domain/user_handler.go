package domain

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetAllUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserHandlerDynamoDb interface {
	Update(c *gin.Context)
	Find(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
}
