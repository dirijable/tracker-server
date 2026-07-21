package repository

import (
	"context"
	"errors"
	"fmt"
	"tracker-system/internal/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) Save(ctx context.Context, user model.User) (model.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return model.User{}, fmt.Errorf("user with tg chat id %q already exist", user.TGChatID)
		}
		return model.User{}, fmt.Errorf("repo save user: %w", err)
	}
	return user, nil
}

func (r *UserRepo) FindByTGChatID(ctx context.Context, tgChatID int64) (model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where(`tg_chat_id = ?`, tgChatID).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, err
		}
		return model.User{}, fmt.Errorf("find by tgChatID: %w", err)
	}
	return user, nil
}
