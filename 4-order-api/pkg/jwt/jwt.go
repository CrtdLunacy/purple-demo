package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

type TokenData struct {
	Phone     string
	SessionId string
	Code      int
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(tokenData *TokenData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone":      tokenData.Phone,
		"session_id": tokenData.SessionId,
		"code":       tokenData.Code,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})
	s, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return s, nil
}
