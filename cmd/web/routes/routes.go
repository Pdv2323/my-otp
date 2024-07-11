package routes

import (
	"net/http"

	"github.com/Pdv2323/my-otp/pkg/config"
	"github.com/Pdv2323/my-otp/pkg/handler"
	"github.com/gin-gonic/gin"
)

func routes(app *config.AppConfig) http.Handler {
	r := gin.Default()
	r.GET("/", handler.FirstHandler)
	r.POST("/verify-otp", handler.FirstHandler)

	return r
}
