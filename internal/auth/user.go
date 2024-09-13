package auth

import (
	"strings"

	"github.com/chscz/backend-onboard-task/internal/config"
)

type UserAuth struct {
	JWTSecretKey     string
	JWTExpiredMinute int
}

func NewUserAuth(jwt config.JWT) *UserAuth {
	return &UserAuth{
		JWTSecretKey:     jwt.SecretKey,
		JWTExpiredMinute: jwt.ExpiredMinute,
	}
}

func (ua *UserAuth) IsValidEmail(email string) bool {
	// emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return strings.Contains(email, "@")
}
