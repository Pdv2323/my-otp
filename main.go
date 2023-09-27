package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pdv2323/otp/temp"
)

func main() {
	r := gin.Default()

	r.GET("/home", temp.Print)

	r.Run(":8888")
}
