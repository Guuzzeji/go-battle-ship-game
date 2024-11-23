package main

import (
	"go-battle-ship-game/routes"

	"github.com/gin-gonic/gin"
)

var server = gin.Default()

func main() {
	router := server.Group("")
	routes.CreateAPIGroup(router)

	server.Run(":8080")
}
