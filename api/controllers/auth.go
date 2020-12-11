package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/elton/my-blog-api/api/models"
	"github.com/elton/my-blog-api/api/responses"
	"github.com/gin-gonic/gin"
)

// Login User Login
func (s *Server) Login(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}

	oneDay, _ := strconv.ParseInt(os.Getenv("OneDayOfHours"), 10, 64)
	audience, _ := jwt.ParseClaimStrings(user.Username)

	// Create the Claims
	claims := &jwt.StandardClaims{
		Audience:  audience,                                         // 受众
		ExpiresAt: jwt.NewTime(float64(time.Now().Unix() + oneDay)), // 失效时间
		IssuedAt:  jwt.NewTime(float64(time.Now().Unix())),          // 签发时间
		Issuer:    "zhiyuan.fun",                                    // 签发人
		NotBefore: jwt.NewTime(float64(time.Now().Unix())),          // 生效时间
		Subject:   "login",                                          // 主题
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("API_SECRET") + user.Password))
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
	}

	responses.ResultJSON(ctx, http.StatusOK, ss, nil)
}

func (s *Server) Register(ctx *gin.Context) {

}
