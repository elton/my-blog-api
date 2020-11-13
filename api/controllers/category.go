package controllers

import (
	"net/http"

	"github.com/elton/my-blog-api/api/models"
	"github.com/gin-gonic/gin"
)

// CreateCategory create a new category
func (s *Server) CreateCategory(ctx *gin.Context) {
	category := models.Category{}
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": nil, "error": err.Error()})
		return
	}
	cateCreated, err := category.SaveCategory(s.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "data": nil, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": cateCreated, "error": nil})
}
