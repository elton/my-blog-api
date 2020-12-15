package controllers

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/elton/my-blog-api/api/models"
	"github.com/elton/my-blog-api/api/responses"
	"github.com/elton/my-blog-api/api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login User Login
func (s *Server) Login(ctx *gin.Context) {
	loginUser := models.User{}
	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}

	users, err := loginUser.FindUsersBy(s.DB)
	if err != gorm.ErrRecordNotFound && err != nil {
		responses.ResultJSON(ctx, http.StatusBadRequest, nil, err)
		return
	}

	if len(*users) > 0 { // 找到了注册用户
		user := (*users)[0]

		if !utils.ComparePasswords(user.Password, []byte(loginUser.Password)) {
			err = errors.New("password is incorrect")
			responses.ResultJSON(ctx, http.StatusOK, nil, err)
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
	} else { // 没有找到就意味着登录的这个用户不存在或者用户名密码错误
		err = errors.New("Invalid user credentials")
		responses.ResultJSON(ctx, http.StatusOK, nil, err)
		return
	}
}

// Register registers a new user
func (s *Server) Register(ctx *gin.Context) {

}
