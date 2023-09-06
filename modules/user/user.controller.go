package userModule

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/abdelrhman-basyoni/gobooks/config"
	domain "github.com/abdelrhman-basyoni/gobooks/core/domain/useCases"
	mongoRepos "github.com/abdelrhman-basyoni/gobooks/core/implementation/repositories/mongo"
	"github.com/abdelrhman-basyoni/gobooks/dto"
	customErrors "github.com/abdelrhman-basyoni/gobooks/errors"
	"github.com/abdelrhman-basyoni/gobooks/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate = validator.New()

func CreateUser(ginContext *gin.Context) {

	var user models.User

	// validate the request body
	if err := ginContext.BindJSON(&user); err != nil {
		ginContext.Error(err)

		return

	}

	// use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		ginContext.Error(validationErr)

		return

	}
	userRepo := mongoRepos.UserRepo{}

	err := domain.CreateUser(user.UserName, user.Password, user.Email, &userRepo)
	if err != nil {
		customErr := &customErrors.DataBaseError{
			Message: err.Error(),
		}

		ginContext.Error(customErr)

		return
	}

	ginContext.JSON(http.StatusCreated, dto.UserResponse{Data: map[string]interface{}{"success": true}})
}

func GetUser(ginContext *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := ginContext.Param("id")
	var user models.User
	defer cancel()
	// change the id from string to ObjectID
	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, dto.UserResponse{Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ginContext.JSON(http.StatusOK, dto.UserResponse{Data: map[string]interface{}{"data": user}})
}
func GetUsers(ginContext *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Print(err.Error())
		ginContext.JSON(http.StatusInternalServerError, dto.UserResponse{Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	//reading from the db in an optimal way
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleUser models.User
		fmt.Printf(singleUser.UserName)
		if err = results.Decode(&singleUser); err != nil {
			ginContext.JSON(http.StatusInternalServerError, dto.UserResponse{Data: map[string]interface{}{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}

	ginContext.JSON(http.StatusOK, dto.UserResponse{Data: map[string]interface{}{"data": users}})
}

func EditUser(ginContext *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := ginContext.Param("id")
	var user models.User
	var updatedUser models.User

	objId, _ := primitive.ObjectIDFromHex(userId)

	defer cancel()
	if err := ginContext.BindJSON(&user); err != nil {
		ginContext.JSON(http.StatusBadRequest, dto.UserResponse{Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	update := bson.M{}
	userdata := reflect.ValueOf(user)
	usertypes := userdata.Type()
	for i := 0; i < userdata.NumField(); i++ {

		if usertypes.Field(i).Name == "Id" {
			fmt.Printf("failed")
			continue
		}
		fieldvalue := userdata.Field(i).Interface().(string)

		if fieldvalue != "" {
			fmt.Println(strings.ToLower(usertypes.Field(i).Name), fieldvalue)
			update = bson.M{strings.ToLower(usertypes.Field(i).Name): fieldvalue}
		}

	}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	err := userCollection.FindOneAndUpdate(ctx, bson.M{"_id": objId}, bson.M{"$set": update}, &opt).Decode(&updatedUser)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, dto.UserResponse{Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	ginContext.JSON(http.StatusOK, dto.UserResponse{Data: map[string]interface{}{"data": updatedUser}})

}
