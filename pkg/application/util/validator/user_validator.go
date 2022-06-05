package validator

import (
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//UserValidator validates user
func UserValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user domain.User
		err := c.ShouldBindJSON(&user)
		if err == nil {
			validate := validator.New()
			err := validate.Struct(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
