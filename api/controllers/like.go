package controllers

import (
	"net/http"

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
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}

	likeGotton, err := like.SaveLikes(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusCreated, likeGotton, nil)
}
