package routes

import (
	"be-golang-chapter-56/user-service/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	userController := controller.NewUserController(db)

	userRoutes := router.Group("/")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
	}
}
