package domain

import (
	domain "github.com/abdelrhman-basyoni/gobooks/core/domain/repositories"
)

func CreateUser(username, password, email string, userRepo domain.UserInterface) error {

	return userRepo.Create(username, password, email)

}
