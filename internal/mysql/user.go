package mysql

import (
	"context"

	"github.com/chscz/backend-onboard-task/internal/domain"
)

func (r UserRepo) CreateUser(ctx context.Context, user *domain.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r UserRepo) GetUser(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	if err := r.DB.WithContext(ctx).
		Model(domain.User{}).
		Where("email = ?", email).
		Take(&user).
		Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}
