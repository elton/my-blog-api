package controllers

import (
	"net/http"
	"strconv"

	"github.com/elton/my-blog-api/api/models"
	"github.com/elton/my-blog-api/api/responses"

	"github.com/gin-gonic/gin"
)

// CreatePost creates a new post.
// curl -i -X POST \
//   http://127.0.0.1:8080/api/v1/posts/\?cid\=1\&uid\=1 \
//   -H 'cache-control: no-cache' \
//   -H 'content-type: application/json' \
//   -d '{
//         "title":"A Post.","summary":"summary of the post","content":"some content of the post."
// }'
func (s *Server) CreatePost(ctx *gin.Context) {
	post := models.Post{}
	if err := ctx.ShouldBindJSON(&post); err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}

	userID, err := strconv.ParseUint(ctx.Query("uid"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	categoryID, err := strconv.ParseUint(ctx.Query("cid"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	postGotton, err := post.SavePost(s.DB, userID, categoryID)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusCreated, postGotton, nil)
}

// FindPostsByUser find posts by user.
func (s *Server) FindPostsByUser(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Query("uid"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}

	var post models.Post
	posts, err := post.FindPostsByUser(s.DB, userID)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, posts, nil)
}

// FindPostByID returns a post by specific post ID.
func (s *Server) FindPostByID(ctx *gin.Context) {
	pid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	post := models.Post{ID: pid}
	postGotton, err := post.FindPostByID(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, postGotton, nil)
}

// FindPostsByCategory returns a list of posts by specific category.
func (s *Server) FindPostsByCategory(ctx *gin.Context) {
	cid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	post := models.Post{}
	posts, err := post.FindPostsByCategory(s.DB, cid)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, posts, nil)
}
