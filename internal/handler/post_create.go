package handler

import (
	"context"
	"net/http"

	"github.com/chscz/backend-onboard-task/internal/domain"
	"github.com/gin-gonic/gin"
)

func (ph *PostHandler) CreatePostPage(c *gin.Context) {
	c.HTML(http.StatusOK, "post_create.tmpl", gin.H{
		"title": "게시글 작성하기",
	})
}

func (ph *PostHandler) CreatePost(c *gin.Context) {
	ctx := context.Background()

	title := c.PostForm("title")
	content := c.PostForm("content")

	tokenString, err := c.Cookie("access-token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{
			"code":    http.StatusUnauthorized,
			"message": "JWT 토큰이 없음",
		})
		return
	}

	userID, err := ph.auth.GetUserIDFromJWT(tokenString)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{
			"code":    http.StatusUnauthorized,
			"message": "토큰이 유효하지 않음",
		})
		return
	}

	post := &domain.Post{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	if err := ph.repo.CreatePost(ctx, post); err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "게시글 작성 실패",
			"err_msg": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
