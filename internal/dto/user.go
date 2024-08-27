package dto

import "github.com/lite07/lite-saas/internal/models"

type UserDto struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func UserDtoFromEntity(user models.User) UserDto {
	return UserDto{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.FirstName,
		Email:     user.Email,
	}
}

func UserDtoFromEntites(users []models.User) []UserDto {
	var userDtos []UserDto

	for _, user := range users {
		var dtoToAdd = UserDto{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.FirstName,
			Email:     user.Email,
		}
		userDtos = append(userDtos, dtoToAdd)
	}

	return userDtos
}
