package usecase

import (
	"amazon-go/domain"

	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) CreateUser(name, phoneNumber, gender string) (string, error) {
	id, _ := uuid.NewUUID()
	user := &domain.User{
		UserId:      id.String(),
		UserName:    name,
		PhoneNumber: phoneNumber,
		Gender:      gender,
	}

	err := uc.userRepo.Create(user)
	if err != nil {
		return "", err
	}
	return user.UserId, nil
}
