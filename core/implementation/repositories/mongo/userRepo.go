package mongoRepos

import (
	"context"
	"fmt"
	"time"

	"github.com/abdelrhman-basyoni/gobooks/config"
	"github.com/abdelrhman-basyoni/gobooks/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")

type UserRepo struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	UserName string             `bson:"username" json:"username"  validate:"required"`
	Password string             `bson:"password" json:"password" validate:"required"`
	Email    string             `bson:"email" json:"email"   validate:"required,email"`
}

func (u *UserRepo) Create(username, password, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	u.Email = email
	u.Password = password
	u.UserName = username

	fmt.Println("createdUser")
	_, err := userCollection.InsertOne(ctx, u)
	return err
}

func (u *UserRepo) GetUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	objId, _ := primitive.ObjectIDFromHex(id)
	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var users []models.User
	results, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleUser models.User

		if err = results.Decode(&singleUser); err != nil {
			return users, err
		}

		users = append(users, singleUser)
	}

	return users, nil
}
