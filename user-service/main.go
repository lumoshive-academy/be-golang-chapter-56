package main

import (
	"be-golang-chapter-56/user-service/models"
	"be-golang-chapter-56/user-service/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.User{})

	router := gin.Default()

	routes.UserRoutes(router, db)

	router.Run(":8082")
}
