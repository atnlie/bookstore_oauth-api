package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp () {
	/*
	oauth use localhost:8081 and user-api use localhost:8080
	 */
	routesMap()
	fmt.Println("Token Access Service started...")
	router.Run(":8081")
}
