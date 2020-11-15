package controllers

import (
	"errors"
	"net/http"
	"strconv"

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
		responses.ResultJSON(ctx, http.StatusUnprocessableEntity, nil, err)
		return
	}

	cateCreated, err := category.SaveCategory(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusCreated, cateCreated, nil)
}

// FindCategoryByID return a category by ID.
func (s *Server) FindCategoryByID(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	category := models.Category{}
	categoryGotten, err := category.FindCategoryByID(s.DB, id)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, categoryGotten, nil)
}

// FindCategoryByName returns a list of categories by name.
func (s *Server) FindCategoryByName(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, errors.New("missing name parameter"))
		return
	}
	category := models.Category{}
	categoriesGotten, err := category.FindCategoriesByName(s.DB, name)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, categoriesGotten, nil)
}

// FindCategories returns the first 100 categories in database.
func (s *Server) FindCategories(ctx *gin.Context) {
	category := models.Category{}
	categoriesGotten, err := category.FindCategories(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, categoriesGotten, nil)
}

// UpdateCategory updates the category
// curl -i -X PUT \
//   http://127.0.0.1:8080/api/v1/categories \
//   -H 'cache-control: no-cache' \
//   -H 'content-type: application/json' \
//   -d '{
//         "id":2,
// 		   "name":"mysql server5"
// }'
func (s *Server) UpdateCategory(ctx *gin.Context) {
	category := models.Category{}

	if err := ctx.ShouldBindJSON(&category); err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}

	if err := category.Validate(); err != nil {
		responses.ResultJSON(ctx, http.StatusUnprocessableEntity, nil, err)
		return
	}

	if err := category.UpdateCategory(s.DB); err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, category, nil)
}

// DeleteCategory Delete a category.
func (s *Server) DeleteCategory(ctx *gin.Context) {
	category := models.Category{}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}

	category.ID = id

	if err := category.Delete(s.DB); err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusNoContent, nil, nil)
}
