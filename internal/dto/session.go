package dto

import (
	"time"

	"github.com/lite07/lite-saas/internal/models"
)

type SessionDto struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func SessionDtoFromEntity(session models.Session) SessionDto {
	return SessionDto{
		Token:     session.Token,
		ExpiredAt: session.ExpiredAt,
	}
}
