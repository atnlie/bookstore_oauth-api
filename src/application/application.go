package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp () {
	routesMap()
	fmt.Println("Token Access Service started...")
	router.Run(":8080")
}
