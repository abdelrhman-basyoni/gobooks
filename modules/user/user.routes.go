package userModule

import "github.com/gin-gonic/gin"

type UserRoutes struct {
	UsrGroup *gin.RouterGroup
}

func RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	userRoutes := &UserRoutes{}
	//All routes related to users comes here
	routerGroup.POST("", userRoutes.CreateUser)

}
