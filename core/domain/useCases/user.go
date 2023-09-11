package domain

import (
	"errors"

	domain "github.com/abdelrhman-basyoni/gobooks/core/domain/repositories"
	"github.com/abdelrhman-basyoni/gobooks/models"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCases struct {
	userRepo domain.UserRepository
}

func (uuc *UserUseCases) CreateUser(username, password, email string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	return uuc.userRepo.Create(username, string(hashedPassword), email)

}

func (uuc *UserUseCases) GetUserById(id string) (*models.User, error) {
	return uuc.userRepo.GetUserById(id)
}
func (uuc *UserUseCases) GetAllUsers() ([]models.User, error) {
	return uuc.userRepo.GetAllUsers()
}

func (uuc *UserUseCases) EditUser(id string, update map[string]interface{}) (*models.User, error) {
	return uuc.userRepo.EditUser(id, update)
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCases {
	return &UserUseCases{
		userRepo: repo,
	}
}

func (uuc *UserUseCases) Login(email string, password string) (string, error) {
	candidate, err := uuc.userRepo.GetUserByEmail(email)

	if err != nil {
		return "", errors.New("invalid User Credentials")
	}
	err = candidate.ValidatePassword(password)
	if err != nil {
		return "", errors.New("invalid User Credentials")
	}
	return candidate.SignToken()
}
