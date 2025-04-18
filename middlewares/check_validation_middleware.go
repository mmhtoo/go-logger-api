package middlewares

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mmhtoo/go-logger-api/helpers"
)

func CheckValidationMiddleware(schema any) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := reflect.New(reflect.TypeOf(schema)).Interface()
		if err := c.ShouldBindJSON(obj); err != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest, 
				helpers.NewAPIErrorResponse(err.Error(), "Validation failed!"),
			)
			return
		}
		c.Set("payload", obj)
		c.Next()
	}
}

func CheckQueryValidationMiddleware(schema any) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		obj := reflect.New(reflect.TypeOf(schema)).Interface()
		if err := ctx.ShouldBindQuery(obj); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, 
				helpers.NewAPIErrorResponse(err.Error(), "Validation failed!"),
			)
			return
		}
		ctx.Set("query", obj)
		ctx.Next()
	}
}
