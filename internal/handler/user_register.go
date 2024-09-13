package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/chscz/backend-onboard-task/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

func (uh *UserHandler) RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user_register.tmpl", gin.H{
		"title":   "회원가입",
		"message": c.Query("message"),
	})
}

func (uh *UserHandler) Register(c *gin.Context) {
	ctx := context.Background()
	email := c.PostForm("email")
	password := c.PostForm("password")
	passwordConfirm := c.PostForm("password_confirm")
	name := c.PostForm("name")

	// 이메일 검사
	if !uh.Auth.IsValidEmail(email) {
		c.Redirect(http.StatusFound, "/register?message=이메일에 '@' 미포함")
		return
	}

	// 비밀번호 검사
	if !uh.Auth.IsValidPassword(password) {
		c.Redirect(http.StatusFound, "/register?message=비밀번호가 8자 미만")
		return
	}
	if password != passwordConfirm {
		c.Redirect(http.StatusFound, "/register?message=비밀번호 확인과 불일치")
		return
	}

	// 기존 유저 이메일 중복 여부 검사
	u, err := uh.repo.GetUser(ctx, email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.Redirect(http.StatusFound, "/register?message=내부 서버 에러")
		return
	}
	if u.Email != "" {
		c.Redirect(http.StatusFound, "/register?message=이미 존재하는 이메일")
		return
	}

	// 비밀번호 암호화
	hashPassword, err := uh.Auth.MakeHashPassword(password)
	if err != nil {
		c.Redirect(http.StatusFound, "/register?message=비밀번호 암호화 오류")
		return
	}

	user := &domain.User{
		ID:       shortid.MustGenerate(),
		Email:    email,
		Name:     name,
		Password: hashPassword,
	}
	// db 저장
	if err = uh.repo.CreateUser(ctx, user); err != nil {
		c.Redirect(http.StatusFound, "/register?message=계정 생성 오류")
		return
	}

	c.Redirect(http.StatusFound, "/login?message=계정 생성 완료")
}
