package userModule

import (
	"context"
	"net/http"
	"time"

	"github.com/abdelrhman-basyoni/gobooks/config"
	"github.com/abdelrhman-basyoni/gobooks/dto"
	"github.com/abdelrhman-basyoni/gobooks/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate = validator.New()

func (userroutes *UserRoutes) CreateUser(ginContext *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err := ginContext.BindJSON(&user); err != nil {
		ginContext.JSON(http.StatusBadRequest, dto.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	// use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		ginContext.JSON(http.StatusBadRequest, dto.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, dto.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ginContext.JSON(http.StatusCreated, dto.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}
