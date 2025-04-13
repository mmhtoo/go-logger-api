package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmhtoo/go-logger-api/helpers"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
	 defer func() {
		if err := recover(); err != nil {
		 // Log the error
		 fmt.Println("Recovered from panic:", err)
 
		 // Return an error response to the client
		 c.JSON(http.StatusInternalServerError, helpers.NewAPIErrorResponse("Internal Server Error!", "Failed to serve!"))
		}
	 }()
 
	 c.Next()
	}
 }