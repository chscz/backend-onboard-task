package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ph *PostHandler) GetPostDetail(c *gin.Context) {
	ctx := context.Background()
	paramID := c.Param("id")
	id, _ := strconv.Atoi(paramID)
	post, err := ph.repo.GetPost(ctx, id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "게시글 보기 실패",
			"err_msg": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "post_detail.tmpl", gin.H{
		"title": "게시글 보기",
		"post":  convertFromDomainPost(&post),
	})
}
