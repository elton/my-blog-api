package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/elton/my-blog-api/api/models"
	"github.com/elton/my-blog-api/api/responses"

	"github.com/gin-gonic/gin"
)

// CreatePost creates a new post.
// curl -i -X POST \
//   http://127.0.0.1:8080/api/v1/posts/\?cid\=1,2\&uid\=1 \
//   -H 'cache-control: no-cache' \
//   -H 'content-type: application/json' \
//   -d '{
//         "title":"A Post IV.","summary":"summary of the post","content":"some content of the post."
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

	// 将逗号分隔的字符串参数转化为[]uint64切片
	cidstrs := strings.Split(ctx.Query("cid"), ",")
	var cids []uint64
	for _, v := range cidstrs {
		id, _ := strconv.Atoi(v)
		cids = append(cids, uint64(id))
	}

	postGotton, err := post.SavePost(s.DB, userID, cids)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusCreated, postGotton, nil)
}

// FindPostsBy find posts by user.
func (s *Server) FindPostsBy(ctx *gin.Context) {
	var (
		post  models.Post
		posts *[]models.Post
		err   error
	)
	if ctx.Query("uid") != "" {
		uid, err := strconv.ParseUint(ctx.Query("uid"), 10, 64)
		if err != nil {
			responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
			return
		}

		posts, err = post.FindPostsByUser(s.DB, uid)
		if err != nil {
			responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
			return
		}
	}

	if ctx.Query("cid") != "" {
		cid, err := strconv.ParseUint(ctx.Query("cid"), 10, 64)
		if err != nil {
			responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
			return
		}

		posts, err = post.FindPostsByCategory(s.DB, cid)
		if err != nil {
			responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
			return
		}
	}

	if ctx.Query("title") != "" {
		posts, err = post.FindPostsByTitle(s.DB, ctx.Query("title"))
		if err != nil {
			responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
			return
		}
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

// FindPosts returns a list of posts.
func (s *Server) FindPosts(ctx *gin.Context) {
	post := models.Post{}
	posts, err := post.FindPosts(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, posts, nil)
}
