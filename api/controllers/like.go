package controllers

import (
	"net/http"
	"strconv"

	"github.com/elton/my-blog-api/api/models"
	"github.com/elton/my-blog-api/api/responses"
	"github.com/gin-gonic/gin"
)

// curl -i -X POST \
//   http://127.0.0.1:8080/api/v1/likes/ \
//   -H 'cache-control: no-cache' \
//   -H 'content-type: application/json' \
//   -d '{
//         "post_id":1,"user_id":1
// }'

// CreateLike creates a new like.
func (s *Server) CreateLike(ctx *gin.Context) {
	var like *models.Like
	if err := ctx.ShouldBindJSON(&like); err != nil {
		responses.ResultJSON(ctx, http.StatusUnprocessableEntity, nil, err)
		return
	}

	likeGotton, err := like.SaveLikes(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusCreated, likeGotton, nil)
}

// FindLikeByID returns the Like object by its ID.
func (s *Server) FindLikeByID(ctx *gin.Context) {
	lid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	like := &models.Like{ID: lid}
	likeGotton, err := like.FindLikeByID(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, likeGotton, nil)
}

// FindLikesBy returns a list of Like objects by given user id or post id.
// curl -i http://localhost:8080/api/v1/likes/\?pid\=1
func (s *Server) FindLikesBy(ctx *gin.Context) {
	var like models.Like
	if ctx.Query("uid") != "" {
		uid, err := strconv.ParseUint(ctx.Query("uid"), 10, 64)
		if err != nil {
			responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
			return
		}
		like.UserID = uid
	}

	if ctx.Query("pid") != "" {
		pid, err := strconv.ParseUint(ctx.Query("pid"), 10, 64)
		if err != nil {
			responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
			return
		}
		like.PostID = pid
	}
	likes, err := like.FindLikesBy(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, likes, nil)
}

// UpdateLike updates the given Like object
func (s *Server) UpdateLike(ctx *gin.Context) {
	lid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	var like models.Like
	if err := ctx.ShouldBindJSON(&like); err != nil {
		responses.ResultJSON(ctx, http.StatusUnprocessableEntity, nil, err)
		return
	}
	like.ID = lid
	if err := like.UpdateLike(s.DB); err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, like, nil)
}

// DeleteLike deletes the given Like object by its ID.
func (s *Server) DeleteLike(ctx *gin.Context) {
	lid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	like := models.Like{ID: lid}
	if err := like.DeleteLike(s.DB); err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusNoContent, nil, nil)
}
