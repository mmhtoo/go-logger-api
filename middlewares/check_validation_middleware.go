package middlewares

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func CheckValidationMiddleware(schema interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := reflect.New(reflect.TypeOf(schema)).Interface()
		if err := c.ShouldBindJSON(obj); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Set("body", obj)
		c.Next()
	}
}
