package sqlRepos

import (
	"fmt"

	"github.com/abdelrhman-basyoni/gobooks/models"
)

type UserRepo struct {
	Id       int    `bson:"_id" json:"id"`
	UserName string `bson:"username" json:"username"  validate:"required"`
	Password string `bson:"password" json:"password" validate:"required"`
	Email    string `bson:"email" json:"email"   validate:"required,email"`
}

func (u *UserRepo) Create(username, password, email string) error {
	fmt.Println("createdUser")
	return nil
}

func (u *UserRepo) GetUserById(id string) (*models.User, error) {
	fmt.Println("createdUser")
	return nil, nil
}
