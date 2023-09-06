package domain

import (
	domain "github.com/abdelrhman-basyoni/gobooks/core/domain/repositories"
	"github.com/abdelrhman-basyoni/gobooks/models"
)

type UserUseCases struct {
	userRepo domain.UserInterface
}

func (uuc *UserUseCases) New(repo domain.UserInterface) {
	uuc.userRepo = repo
}
func (uuc *UserUseCases) CreateUser(username, password, email string) error {

	return uuc.userRepo.Create(username, password, email)

}

func (uuc *UserUseCases) GetUserById(id string) (*models.User, error) {
	return uuc.userRepo.GetUserById(id)
}
