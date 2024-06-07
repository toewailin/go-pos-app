package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/toewailin/pos-app/controllers"
	"github.com/toewailin/pos-app/middleware"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/login", controllers.Login)
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		r.GET("/items", controllers.GetItems(db))
		r.POST("/items", controllers.CreateItem(db))
	}
	return r
}
