package temp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Print(c *gin.Context) {
	c.String(http.StatusOK, "This is OTP generator code.\n")
	c.String(http.StatusOK, "Let's get started with this.")
}
