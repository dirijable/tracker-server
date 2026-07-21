package dto

import "github.com/google/uuid"

type UserResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	TGChatID int64     `json:"tg_chat_id"`
}
