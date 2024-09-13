package handler

import (
	"time"

	"github.com/chscz/backend-onboard-task/internal/domain"
)

type Post struct {
	ID        int
	CreatedAt string
	UpdatedAt string
	UserID    string
	Title     string
	Content   string
	ViewCount int
	UserName  string
}

func (p Post) convertToDomain() (*domain.Post, error) {
	createdAt, _ := time.Parse("2006-01-02 15:04:05", p.CreatedAt)
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", p.UpdatedAt)
	return &domain.Post{
		ID:        p.ID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		ViewCount: p.ViewCount,
	}, nil
}

func convertFromDomainPostList(posts []*domain.Post) []*Post {
	postList := make([]*Post, len(posts))
	for i, post := range posts {
		postList[i] = convertFromDomainPost(post)
	}
	return postList
}

func convertFromDomainPost(p *domain.Post) *Post {
	return &Post{
		ID:        p.ID,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		ViewCount: p.ViewCount,
		UserName:  p.User.Name,
	}
}

func getFirstLastPostID(items []*domain.Post) (first, last int) {
	first = int(items[0].ID)
	last = int(items[len(items)-1].ID)
	return
}

func setPostPage(page, totalCount int) (prevPage, nextPage int) {
	// 이전 페이지와 다음 페이지를 계산
	prevPage = page - 1
	nextPage = page + 1

	// 이전 페이지가 1보다 작으면 이전 페이지는 1로 설정
	if prevPage < 1 {
		prevPage = 1
	}

	// 다음 페이지가 마지막 페이지를 넘어가면 다음 페이지는 마지막 페이지로 설정
	if nextPage > (totalCount+itemsPerPage-1)/itemsPerPage {
		nextPage = (totalCount + itemsPerPage - 1) / itemsPerPage
	}
	return
}
