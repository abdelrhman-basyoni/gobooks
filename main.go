package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	/**
	create router and listen on a a port
	*/
	router := gin.Default()
	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusAccepted, "it works")
	})
	fmt.Printf("server runnung on port 6000 \n")
	log.Fatalln(router.Run(":6000"))

}
