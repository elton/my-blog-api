package controllers

import (
	"net/http"

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
	// pid, err := strconv.ParseUint(ctx.Query("pid"), 10, 64)
	// if err != nil {
	// 	responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
	// 	return
	// }
	// uid, err := strconv.ParseUint(ctx.Query("uid"), 10, 64)
	// if err != nil {
	// 	responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
	// 	return
	// }
	// comment.PostID = pid
	// comment.UserID = uid
	commentGotton, err := comment.SaveComment(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusCreated, commentGotton, nil)
}
