package domain

import "github.com/gin-gonic/gin"

//AuthHandler Auth handler
type AuthHandler interface {
	Login(c *gin.Context)
	Refresh(c *gin.Context)
	Logout(c *gin.Context)
}
