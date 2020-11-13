package controllers

import (
	"net/http"

	"github.com/elton/my-blog-api/api/models"
	"github.com/elton/my-blog-api/api/responses"
	"github.com/gin-gonic/gin"
)

// CreateCategory create a new category
func (s *Server) CreateCategory(ctx *gin.Context) {
	category := models.Category{}
	if err := ctx.ShouldBindJSON(&category); err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}

	if err := category.Validate(); err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}

	cateCreated, err := category.SaveCategory(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusCreated, cateCreated, nil)
}
