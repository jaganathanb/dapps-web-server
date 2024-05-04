package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, "Something went wrong!")
}
