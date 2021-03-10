package application

import ping "amiera/src/controllers/controller_ping"

func routesMap () {
	router.GET("/ping", ping.Ping)
}
