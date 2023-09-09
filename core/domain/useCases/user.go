package domain

import (
	domain "github.com/abdelrhman-basyoni/gobooks/core/domain/repositories"
	"github.com/abdelrhman-basyoni/gobooks/models"
)

type UserUseCases struct {
	userRepo domain.UserRepository
}

func (uuc *UserUseCases) CreateUser(username, password, email string) error {

	return uuc.userRepo.Create(username, password, email)

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
