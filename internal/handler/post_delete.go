package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ph *PostHandler) DeletePost(c *gin.Context) {
	defer c.Redirect(http.StatusFound, "/")

	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	post, err := ph.repo.GetPost(ctx, id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "삭제할 게시글 조회 실패",
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

	if !ph.MatchUserID(post.UserID, tokenString) {
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{
			"code":    http.StatusUnauthorized,
			"message": "작성자와 불일치",
		})
	}

	if err := ph.repo.DeletePost(ctx, id); err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "게시글 삭제 실패",
			"err_msg": err.Error(),
		})
		return
	}
}
