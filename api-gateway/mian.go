package main

import (
	"be-golang-chapter-56/api-gateway/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middlewares
	router.Use(middleware.AuthMiddleware())

	// Routes
	router.Any("/product/*proxyPath", reverseProxy("http://localhost:8081"))
	router.Any("/user/*proxyPath", reverseProxy("http://localhost:8082"))

	router.Run(":8080")
}

func reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyPath := c.Param("proxyPath")
		targetURL := target + proxyPath
		http.Redirect(c.Writer, c.Request, targetURL, http.StatusTemporaryRedirect)
	}
}
