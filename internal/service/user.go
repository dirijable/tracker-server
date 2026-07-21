package service

import (
	"context"
	"fmt"
	"tracker-system/internal/dto"
	"tracker-system/internal/model"
	"tracker-system/internal/service/mapper"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User) (model.User, error)
	FindByTGChatID(ctx context.Context, tgChatID int64) (model.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Save(ctx context.Context, tgChatID int64) error {
	user := model.User{TGChatID: tgChatID}
	_, err := u.repo.Save(ctx, user)
	if err != nil {
		return fmt.Errorf("save user: %w", err)
	}
	return nil
}

func (u *UserService) FindByTGChatID(ctx context.Context, tgChatID int64) (dto.UserResponseDTO, error) {
	user, err := u.repo.FindByTGChatID(ctx, tgChatID)
	if err != nil {
		return dto.UserResponseDTO{}, fmt.Errorf("find by tg chat id: %w", err)
	}
	return mapper.UserModelToResponse(user), nil
}
