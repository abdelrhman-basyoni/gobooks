package domain

import "github.com/abdelrhman-basyoni/gobooks/models"

type UserRepository interface {
	Create(username, password, email string) error
	GetUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	EditUser(id string, update map[string]interface{}) error
}
