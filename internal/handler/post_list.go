package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chscz/backend-onboard-task/internal/domain"
	"github.com/gin-gonic/gin"
)

func (ph *PostHandler) GetPosts(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	// 페이징처리 모드
	mode := pageModeHome
	if c.Query("mode") == "next" {
		mode = pageModeNext
	} else if c.Query("mode") == "prev" {
		mode = pageModePrev
	}

	ctx := context.Background()
	posts := make([]*domain.Post, 0)

	switch mode {
	case pageModeHome:
		posts, err = ph.repo.GetPosts(ctx, itemsPerPage, "")
	case pageModeNext:
		cursor, _ := strconv.Atoi(c.Query("cursor"))
		posts, err = ph.repo.GetPosts(ctx, itemsPerPage, fmt.Sprintf("id < %d", cursor))
	case pageModePrev:
		cursor, _ := strconv.Atoi(c.Query("cursor"))
		posts, err = ph.repo.GetPosts(ctx, 0, fmt.Sprintf("id > %d", cursor))
		posts = posts[len(posts)-itemsPerPage:]
	}
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "목록 조회 실패",
			"err_msg": err.Error(),
		})
		return
	}

	if len(posts) == 0 {
		c.HTML(http.StatusOK, "post_list.tmpl", gin.H{
			"title": "게시글 목록",
		})
		return
	}
	firstItemID, lastItemID := getFirstLastPostID(posts)

	// 페이징 정보 계산
	totalCount, err := ph.repo.GetTotalPostCount(ctx)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "전체 게시글 수 조회 실패",
			"err_msg": err.Error(),
		})
		return
	}
	prevPage, nextPage := setPostPage(page, totalCount)

	c.HTML(http.StatusOK, "post_list.tmpl", gin.H{
		"title":          "게시글 목록",
		"posts":          convertFromDomainPostList(posts),
		"firstItemID":    firstItemID,
		"lastItemID":     lastItemID,
		"currentPage":    page,
		"prevPage":       prevPage,
		"nextPage":       nextPage,
		"totalPages":     (totalCount + itemsPerPage - 1) / itemsPerPage,
		"totalPostCount": totalCount,
	})
}
