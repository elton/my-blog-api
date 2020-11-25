package controllers

import (
	"net/http"
	"strconv"

	"github.com/elton/my-blog-api/api/models"
	"github.com/elton/my-blog-api/api/responses"
	"github.com/gin-gonic/gin"
)

// CreateComment creates a new comment.
func (s *Server) CreateComment(ctx *gin.Context) {
	var comment *models.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	commentGotton, err := comment.SaveComment(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusCreated, commentGotton, nil)
}

// FindCommentByID returns a Comment by ID.
func (s *Server) FindCommentByID(ctx *gin.Context) {
	cid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	comment := models.Comment{ID: cid}
	commentGotton, err := comment.FindCommentByID(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, commentGotton, nil)
}

// FindCommentsBy returns a list of comments by criterias.
func (s *Server) FindCommentsBy(ctx *gin.Context) {
	var comment models.Comment
	if ctx.Query("pid") != "" {
		pid, err := strconv.ParseUint(ctx.Query("pid"), 10, 64)
		if err != nil {
			responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
			return
		}
		comment.PostID = pid

	}

	if ctx.Query("uid") != "" {
		uid, err := strconv.ParseUint(ctx.Query("uid"), 10, 64)
		if err != nil {
			responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
			return
		}
		comment.UserID = uid
	}

	comments, err := comment.FindCommentsBy(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, comments, nil)
}

// UpdateComment updates a comment.
func (s *Server) UpdateComment(ctx *gin.Context) {
	cid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	var comment models.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		responses.ResultJSON(ctx, http.StatusUnprocessableEntity, nil, err)
		return
	}
	comment.ID = cid
	if err := comment.UpdateComment(s.DB); err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, comment, nil)
}

// DeleteComment deletes a comment.
func (s *Server) DeleteComment(ctx *gin.Context) {
	cid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	comment := models.Comment{ID: cid}
	if err := comment.DeleteComment(s.DB); err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusNoContent, nil, nil)
}
