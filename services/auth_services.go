package services

import (
	"gin_fleamarket/models"
	"gin_fleamarket/repository"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, password string) error
}

type AuthService struct {
	repository repository.IAuthRepository
}

func NewAuthService(repository repository.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Signup(email string, password string) error {
	hashdpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashdpassword),
	}
	return s.repository.CreateUser(user)
}
