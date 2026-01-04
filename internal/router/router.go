package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"SM-Go/internal/auth"
	"SM-Go/internal/db"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"status":"ok"})
}

func RoutesHandler(r *gin.Engine) {
	r.GET("/health", healthCheck)
	authRepo:=auth.NewRepo(db.DB)
	authService:=auth.NewService(authRepo)
	authHandler:=auth.NewHandler(authService)
	authGroup:=r.Group("/auth")
	{
		authGroup.POST("/register",authHandler.Register)
		authGroup.POST("/login",authHandler.Login)
	}
}