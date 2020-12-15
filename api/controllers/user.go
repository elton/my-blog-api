package controllers

import (
	"net/http"
	"strconv"

	"github.com/elton/my-blog-api/api/models"
	"github.com/elton/my-blog-api/api/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUser creates a new user.
func (s *Server) CreateUser(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}

	if err := user.Validate(); err != nil {
		responses.ResultJSON(ctx, http.StatusUnprocessableEntity, nil, err)
		return
	}

	userCreated, err := user.SaveUser(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusCreated, userCreated, nil)
}

// FindUserByID returns a user matching specific ID.
func (s *Server) FindUserByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	user := models.User{ID: id}
	userGotton, err := user.FindUserByID(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, userGotton, nil)
}

// FindUsers returns a list of users.
func (s *Server) FindUsers(ctx *gin.Context) {
	user := models.User{}
	usersGotton, err := user.FindUsers(s.DB)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, usersGotton, nil)
}

// FindUsersBy returns a list of users matching the given criterias.
func (s *Server) FindUsersBy(ctx *gin.Context) {
	var (
		t   uint64
		err error
	)
	if ctx.Query("type") != "" {
		t, err = strconv.ParseUint(ctx.Query("type"), 10, 8)
		if err != nil {
			responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
			return
		}
	}

	user := models.User{
		Username: ctx.Query("username"),
		Nickname: ctx.Query("nickname"),
		Type:     uint8(t),
		Mobile:   ctx.Query("mobile"),
		Email:    ctx.Query("email"),
	}
	usersGotten, err := user.FindUsersBy(s.DB)

	switch {
	case err == gorm.ErrRecordNotFound:
		{
			responses.ResultJSON(ctx, http.StatusOK, nil, err)
			return
		}
	case err != nil:
		{
			responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
			return
		}
	default:
		responses.ResultJSON(ctx, http.StatusOK, usersGotten, nil)
	}
}

// UpdateUser  updates a user.
func (s *Server) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	user := models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	user.ID = id

	if err := user.UpdateUser(s.DB); err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusOK, user, nil)
}

// DeleteUser deletes a user
func (s *Server) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}
	user := models.User{}
	user.ID = id
	if err := user.DeleteUser(s.DB); err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	responses.ResultJSON(ctx, http.StatusNoContent, nil, nil)
}
