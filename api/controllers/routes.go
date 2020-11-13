package controllers

import "github.com/gin-gonic/gin"

// Routers routers for the application.
func Routers(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello/:name", Hello)
	}
}
