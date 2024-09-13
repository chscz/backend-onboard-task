package handler

import (
	"context"

	"github.com/chscz/backend-onboard-task/internal/auth"
	"github.com/chscz/backend-onboard-task/internal/domain"
)

type UserHandler struct {
	repo UserRepository
	Auth *auth.UserAuth
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, id string) (domain.User, error)
}

func NewUserHandler(repo UserRepository, auth *auth.UserAuth) *UserHandler {
	return &UserHandler{
		repo: repo,
		Auth: auth,
	}
}
