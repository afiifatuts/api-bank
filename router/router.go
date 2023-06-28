package router

import (
	"github.com/afiifatuts/bankmnc/handler"
	"github.com/afiifatuts/bankmnc/middleware"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.POST("/login", handler.Login())
	router.Use(middleware.AuthenticationMiddleware())
	router.POST("/logout", handler.Logout())
	router.POST("/payment", handler.Payment())

	router.Run("localhost:8000")

	return router
}
