package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/lite07/lite-saas/internal/database"
	"github.com/lite07/lite-saas/internal/dto"
	"github.com/lite07/lite-saas/internal/models"
	"github.com/lite07/lite-saas/internal/requests"
	"github.com/lite07/lite-saas/internal/utils"
)

var ErrorInvalidCredential = errors.New("user or password is incorrect")

func CreateUserSession(request requests.CreateUserSession) (dto.SessionDto, error) {
	var user models.User
	if err := database.DB.Where("Email = ?", request.Email).First(&user).Error; err != nil {
		return dto.SessionDto{}, ErrorInvalidCredential
	}

	var hashedPassword = utils.HashSha256String(request.Password)
	if hashedPassword != user.PasswordHash {
		return dto.SessionDto{}, ErrorInvalidCredential
	}

	sessionToken := models.SessionToken{
		UserEmail: user.Email,
		FullName:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
	}

	sessionTokenJson, err := json.Marshal(sessionToken)
	if err != nil {
		panic(err)
	}

	session := models.Session{
		ExpiredAt: time.Now().Local().Add(time.Minute * time.Duration(60)),
		UserId:    user.ID,
		User:      user,
		Token:     utils.EncryptString(string(sessionTokenJson)),
	}
	if err := database.DB.Create(&session).Error; err != nil {
		panic(err)
	}

	return dto.SessionDtoFromEntity(session), nil
}
