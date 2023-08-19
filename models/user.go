package models

import (
	"context"
	"fmt"

	"github.com/abdelrhman-basyoni/gobooks/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	UserName string             `bson:"username" json:"username"  validate:"required"`
	Password string             `bson:"password" json:"password" validate:"required"`
	Email    string             `bson:"email" json:"email"   validate:"required,email"`
}

func userIndex() string {
	var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")

	// Define the unique index model
	keys := bson.D{{"email", 1}} // Fields you want to create a unique index on
	indexModel := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(true), // SetUnique creates a unique index
	}

	// Create the unique index
	indexName, err := userCollection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		fmt.Println("Error creating index:", err)

	}
	return indexName
}

var indexName = userIndex()
