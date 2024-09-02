package routes

import (
	"be-golang-chapter-56/product-service/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(router *gin.Engine, db *gorm.DB) {
	productController := controller.NewProductController(db)

	productRoutes := router.Group("/product")
	{
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/:id", productController.GetProductByID)
		productRoutes.GET("/", productController.GetAllProducts)
	}
}
