package userModule

import (
	"net/http"

	"github.com/abdelrhman-basyoni/gobooks/config"
	domain "github.com/abdelrhman-basyoni/gobooks/core/domain/useCases"
	mongoRepos "github.com/abdelrhman-basyoni/gobooks/core/implementation/repositories/mongo"
	"github.com/abdelrhman-basyoni/gobooks/dto"
	customErrors "github.com/abdelrhman-basyoni/gobooks/errors"
	"github.com/abdelrhman-basyoni/gobooks/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate = validator.New()
var useCases = domain.NewUserUseCase(&mongoRepos.UserRepo{})

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

	err := useCases.CreateUser(user.UserName, user.Password, user.Email)
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
	userId := ginContext.Param("id")

	user, err := useCases.GetUserById(userId)
	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusOK, dto.UserResponse{Data: user})
}

func GetUsers(ginContext *gin.Context) {

	users, err := useCases.GetAllUsers()

	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusOK, dto.UserResponse{Data: map[string]interface{}{"users": users}})
}

func EditUser(ginContext *gin.Context) {

	userId := ginContext.Param("id")

	var update map[string]interface{}
	if err := ginContext.BindJSON(&update); err != nil {
		ginContext.Error(err)
		return
	}
	updatedUser, err := useCases.EditUser(userId, update)
	if err != nil {
		ginContext.Error(err)
		return
	}
	ginContext.JSON(http.StatusOK, dto.UserResponse{Data: map[string]interface{}{"updatedUser": updatedUser}})

}
