package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (uh *UserHandler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user_login.tmpl", gin.H{
		"title":   "로그인",
		"message": c.Query("message"),
	})
}

func (uh *UserHandler) Login(c *gin.Context) {
	ctx := context.Background()
	email := c.PostForm("email")
	password := c.PostForm("password")

	// 이메일 검사
	if !uh.Auth.IsValidEmail(email) {
		c.Redirect(http.StatusFound, "/login?message=이메일에 '@' 미포함")
		return
	}

	// 비밀번호 검사
	if !uh.Auth.IsValidPassword(password) {
		c.Redirect(http.StatusFound, "/login?message=비밀번호가 8자 미만")
		return
	}

	// 계정 조회
	user, err := uh.repo.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Redirect(http.StatusFound, "/login?message=존재하지 않는 유저")
			return
		}
		c.Redirect(http.StatusFound, "/login?message=서버 내부 에러")
		return
	}

	// 비밀번호 검사
	if !uh.Auth.CheckPasswordHash(user.Password, password) {
		c.Redirect(http.StatusFound, "/login?message=비밀번호 불일치")
		return
	}

	// 토큰 발행
	accessToken, err := uh.Auth.CreateJWT(user.ID)
	if err != nil {
		c.Redirect(http.StatusFound, "/login?message=토큰 발행 오류")
		return
	}

	// 쿠키 생성
	cookie, err := c.Cookie("access-token")
	if err != nil {
		cookie = "NotSet"
		c.SetCookie(
			"access-token",
			accessToken,
			3600,
			"/",
			"localhost",
			false,
			true,
		)
	}
	_ = cookie

	c.Redirect(http.StatusFound, "/")
}
