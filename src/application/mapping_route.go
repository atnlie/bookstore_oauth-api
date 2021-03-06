package application

import (
	"amiera/src/controllers/controller_http"
	ping "amiera/src/controllers/controller_ping"
	"amiera/src/repository/db"
	"amiera/src/service"
)

func routesMap () {
	router.GET("/ping", ping.Ping)
	atService := service.NewService(db.NewRepository())
	atHandler := controller_http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
}
