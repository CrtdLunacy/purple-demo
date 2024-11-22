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
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func generateSerssion(phone string) *VerifyData {
	// создаем локальный рандомизатор
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)

	//генерируем рандомный 4-значный код
	randomNumber := localRand.Intn(9000) + 1000

	sessionId, _ := hash.HashString(phone)

	return &VerifyData{
		SessionId: sessionId,
		Code:      randomNumber,
	}
}

func (service *AuthService) Auth(data *AuthData) (*VerifyData, error) {
	existedUser, _ := service.UserRepository.FindByPhone(data.Phone)

	if existedUser == nil {
		session := generateSerssion(data.Phone)
		_, err := service.UserRepository.Create(&user.User{
			Phone:     data.Phone,
			SessionId: session.SessionId,
			Code:      session.Code,
		})
		if err != nil {
			return nil, errors.New(ErrCommon)
		}

		return session, nil
	}

	session := &VerifyData{
		SessionId: existedUser.SessionId,
		Code:      existedUser.Code,
	}

	return session, nil
}

func (service *AuthService) Verify(data *VerifyData) (*user.User, error) {
	user, err := service.UserRepository.FindBySession(data.SessionId, data.Code)
	if err != nil {
		return nil, errors.New(ErrUnverified)
	}

	return user, nil
}
