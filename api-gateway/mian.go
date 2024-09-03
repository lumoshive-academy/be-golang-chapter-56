package main

import (
	"be-golang-chapter-56/api-gateway/middleware"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// User Routes - No Authentication Required
	router.Any("/user/*proxyPath", reverseProxy("http://localhost:8082", "/user"))

	// Product Routes - Authentication Required
	productRoutes := router.Group("/product")
	productRoutes.Use(middleware.AuthMiddleware())
	productRoutes.Any("/*proxyPath", reverseProxy("http://localhost:8081", "/product"))

	router.Run(":8080")
}

func reverseProxy(target, prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyPath := strings.TrimPrefix(c.Param("proxyPath"), prefix)
		targetURL := target + proxyPath
		http.Redirect(c.Writer, c.Request, targetURL, http.StatusTemporaryRedirect)
	}
}
