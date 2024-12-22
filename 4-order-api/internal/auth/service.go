package auth

import (
	"errors"
	"go/order-api/internal/user"
	"go/order-api/pkg/hash"
	"math/rand"
	"time"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

type AuthData struct {
	Phone string
}

type VerifyData struct {
	SessionId string
	Code      int
	Expiry    time.Time
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func generateSession(phone string) *VerifyData {
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)

	randomNumber := localRand.Intn(9000) + 1000
	sessionId, _ := hash.HashString(phone)

	return &VerifyData{
		SessionId: sessionId,
		Code:      randomNumber,
		Expiry:    time.Now().Add(2 * time.Minute),
	}
}

func (service *AuthService) Auth(data *AuthData) (*VerifyData, error) {
	existedUser, _ := service.UserRepository.FindByPhone(data.Phone)

	if existedUser != nil {
		if time.Now().Before(existedUser.SessionExpiry) {
			// Если сессия ещё действительна, возвращаем её данные
			return &VerifyData{
				SessionId: existedUser.SessionId,
				Code:      existedUser.Code,
				Expiry:    existedUser.SessionExpiry,
			}, nil
		}

		newSession := generateSession(data.Phone)
		existedUser.SessionId = newSession.SessionId
		existedUser.Code = newSession.Code
		existedUser.SessionExpiry = newSession.Expiry

		if err := service.UserRepository.Update(existedUser); err != nil {
			return nil, errors.New(ErrCommon)
		}

		return newSession, nil
	}

	session := generateSession(data.Phone)
	_, err := service.UserRepository.Create(&user.User{
		Phone:         data.Phone,
		SessionId:     session.SessionId,
		Code:          session.Code,
		SessionExpiry: session.Expiry,
	})
	if err != nil {
		return nil, errors.New(ErrCommon)
	}

	return session, nil
}

func (service *AuthService) Verify(data *VerifyData) (*user.User, error) {
	user, err := service.UserRepository.FindBySession(data.SessionId, data.Code)
	if err != nil {
		return nil, errors.New(ErrUnverified)
	}

	if time.Now().After(user.SessionExpiry) {
		return nil, errors.New(ErrSessionExpired)
	}

	return user, nil
}
