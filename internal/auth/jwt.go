package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func (ua *UserAuth) CreateJWT(userID string) (string, error) {
	mySigningKey := []byte(ua.JWTSecretKey)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(ua.JWTExpiredMinute) * time.Minute).Unix(),
		},
	}

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tk, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tk, nil
}

func (ua *UserAuth) GetUserIDFromJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ua.JWTSecretKey), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		userID := claims.UserID
		return userID, nil
	}

	return "", errors.New("invalid token")
}

func (ua *UserAuth) ValidateJWT(jwtKey, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
