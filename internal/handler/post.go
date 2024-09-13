package handler

import (
	"context"

	"github.com/chscz/backend-onboard-task/internal/auth"
	"github.com/chscz/backend-onboard-task/internal/domain"
)

const itemsPerPage = 10

const (
	pageModeHome = iota
	pageModeNext
	pageModePrev
)

const (
	ErrNoToken        = "JWT 토큰이 없음"
	ErrInvalidToken   = "토큰이 유효하지 않음"
	ErrPostCreateFail = "게시글 작성 실패"
)

type PostHandler struct {
	repo PostRepository
	auth *auth.UserAuth
}

type PostRepository interface {
	CreatePost(ctx context.Context, post *domain.Post) error
	UpdatePost(ctx context.Context, post *domain.Post) error
	DeletePost(ctx context.Context, id int) error

	GetPost(ctx context.Context, id int) (domain.Post, error)
	GetPosts(ctx context.Context, itemsPerPage int, whereClause string) ([]*domain.Post, error)
	GetTotalPostCount(ctx context.Context) (int, error)
}

func NewPostHandler(repo PostRepository, auth *auth.UserAuth) *PostHandler {
	return &PostHandler{
		repo: repo,
		auth: auth,
	}
}

func (ph *PostHandler) MatchUserID(userID, tokenString string) bool {
	userIDFromToken, err := ph.auth.GetUserIDFromJWT(tokenString)
	if err != nil {
		return false
	}
	return userID == userIDFromToken
}
