package mysql

import (
	"context"

	"github.com/chscz/backend-onboard-task/internal/domain"
	"gorm.io/gorm"
)

func (r PostRepo) CreatePost(ctx context.Context, post *domain.Post) error {
	return r.DB.WithContext(ctx).Create(post).Error
}

func (r PostRepo) UpdatePost(ctx context.Context, post *domain.Post) error {
	return r.DB.WithContext(ctx).Updates(post).Error
}

func (r PostRepo) DeletePost(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Model(domain.Post{}).Delete("id = ?", id).Error
}

func (r PostRepo) GetPost(ctx context.Context, id int) (domain.Post, error) {
	var post domain.Post
	if err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&domain.Post{}).
			Where("id = ?", id).
			Update("view_count", gorm.Expr("view_count + ?", 1)).
			Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).
			Preload("User").
			Where("id = ?", id).
			Take(&post).
			Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

func (r PostRepo) GetPosts(ctx context.Context, itemsPerPage int, whereClause string) ([]*domain.Post, error) {
	var posts []*domain.Post

	query := r.DB.WithContext(ctx).
		Preload("User").
		Where(whereClause).
		Order("created_at DESC, id DESC")

	if itemsPerPage > 0 {
		query = query.Limit(itemsPerPage)
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r PostRepo) GetTotalPostCount(ctx context.Context) (int, error) {
	var cnt int64
	if err := r.DB.WithContext(ctx).
		Model(domain.Post{}).
		Count(&cnt).
		Error; err != nil {
		return 0, err
	}
	return int(cnt), nil
}
