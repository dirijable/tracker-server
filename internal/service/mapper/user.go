package mapper

import (
	"tracker-system/internal/dto"
	"tracker-system/internal/model"
)

func UserModelToResponse(user model.User) dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID:       user.ID,
		TGChatID: user.TGChatID,
	}
}
