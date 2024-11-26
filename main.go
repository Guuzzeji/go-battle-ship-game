package main

import (
	"go-minesweaper-multiplayer/routes"

	"github.com/gin-gonic/gin"
)

// Creating server default instance
var server = gin.Default()

// Run server
func main() {
	router := server.Group("")
	routes.CreateAPIGroup(router)
	routes.CreateClientGroup(server)

	server.Run(":8080") // Running on localhost:8080
}
