package main

import (
	"fmt"
	"log"
	"net/http"

	ginHandler "github.com/abdelrhman-basyoni/gin-error-handler"
	userModule "github.com/abdelrhman-basyoni/gobooks/modules/user"
	"github.com/gin-gonic/gin"
)

func main() {
	/**
	create router and listen on a a port
	*/

	router := gin.Default()

	router.Use(gin.HandlerFunc(ginHandler.GlobalErrorHandler("errors")))

	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusAccepted, "it works")
	})

	//register the routes

	userModule.RegisterUserRoutes(router.Group(("user")))

	fmt.Printf("server running on port 6000 \n")
	log.Fatalln(router.Run(":6000"))

}
