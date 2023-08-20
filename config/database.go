package config

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/abdelrhman-basyoni/gobooks/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// define the collections
type MongoDB struct {
	client         *mongo.Client
	database       *mongo.Database
	userCollection *mongo.Collection
}

//check if there is an instance of the MongoDB struct

var lock = &sync.Mutex{}
var instance *MongoDB

// func CreateMongoInstance() *MongoDB {
// 	// kinda like singlton
// 	if instance == nil {
// 		lock.Lock()
// 		defer lock.Lock()
// 		if instance == nil {
// 			log.Println("Creating a database instance")
// 			instance = &MongoDB{}
// 			instance.ConnectDB()
// 		}
// 	}
// 	return instance
// }

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(utils.ReadEnv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	//trying to connect to the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	client.Database("golangAPI")
	log.Printf("Connected to MongoDB")
	//register to the instance
	// mongod.client = client
	// mongod.database = client.Database("gobooks")
	// mongod.userCollection = mongod.database.Collection("user")

	return client
}

var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}
