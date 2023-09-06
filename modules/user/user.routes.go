package userModule

import "github.com/gin-gonic/gin"

type UserRoutes struct {
	UsrGroup *gin.RouterGroup
}

func RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	// userRoutes := &UserRoutes{}
	//All routes related to users comes here
	routerGroup.POST("/create", CreateUser)
	routerGroup.GET("/getall/", GetUsers)
	routerGroup.GET("/get/:id", GetUser)
	routerGroup.PUT("/edit/:id", EditUser)
	routerGroup.DELETE("/get/:id", GetUser)

}
