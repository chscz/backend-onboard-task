package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ph *PostHandler) UpdatePostPage(c *gin.Context) {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))
	post, err := ph.repo.GetPost(ctx, id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "수정할 게시글 조회 실패",
			"err_msg": err.Error(),
		})
		return
	}

	tokenString, err := c.Cookie("access-token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{
			"code":    http.StatusUnauthorized,
			"message": "JWT 토큰이 없음",
		})
		return
	}

	if !ph.MatchUserID(post.User.ID, tokenString) {
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{
			"code":    http.StatusUnauthorized,
			"message": "작성자와 불일치",
		})
		return
	}

	c.HTML(http.StatusOK, "post_update.tmpl", gin.H{
		"title": "게시글 수정하기",
		"post":  convertFromDomainPost(&post),
	})
}

func (ph *PostHandler) UpdatePost(c *gin.Context) {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))
	p := &Post{
		ID:      id,
		Title:   c.PostForm("title"),
		Content: c.PostForm("content"),
	}

	post := p.convertToDomain()
	if err := ph.repo.UpdatePost(ctx, post); err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "게시글 수정 실패",
			"err_msg": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}
