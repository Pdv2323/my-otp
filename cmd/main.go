package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/my-otp/pkg/handlers"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("pkg/templates/*")

	indexFile := "index.html"
	verifyFile := "verify.html"

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, indexFile, nil)
	})
	r.POST("/submit-email", handlers.SubmitEmailHandler)

	r.GET("/submit-otp", func(c *gin.Context) {
		c.HTML(http.StatusOK, verifyFile, nil)
	})
	r.POST("/submit-otp", handlers.SubmitOtpHandler)

	r.Run(":1234")
}
