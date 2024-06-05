package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FirstHandler(c *gin.Context) {
	fmt.Println("First-Handler")
}
