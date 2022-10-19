package main

import (
	"fmt"
	"log"
	"net/http"

	userModule "github.com/abdelrhman-basyoni/gobooks/modules/user"
	"github.com/gin-gonic/gin"
)

func main() {
	/**
	create router and listen on a a port
	*/
	// go config.ConnectDB()
	router := gin.Default()
	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusAccepted, "it works")
	})

	//register the routes
	userModule.RegisterUserRoutes(router.Group(("user")))

	fmt.Printf("server runnung on port 6000 \n")
	log.Fatalln(router.Run(":6000"))

}
