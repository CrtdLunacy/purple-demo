package user

import "go/order-api/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepository) FindBySession(sessionId string, code int) (*User, error) {
	var user User
	err := repo.Database.DB.First(&user, "session_id = ? AND code = ?", sessionId, code).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
