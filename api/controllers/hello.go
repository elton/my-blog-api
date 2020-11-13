package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello say hello
func (s *Server) Hello(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{"msg": "Hello" + name})
}
