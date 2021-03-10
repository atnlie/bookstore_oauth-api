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

	fmt.Println("Siap laksanakan")

	router.Run(":8080")
}