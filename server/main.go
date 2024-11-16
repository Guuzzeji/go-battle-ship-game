package main

import (
	api "battleship-server/routes"

	"github.com/gin-gonic/gin"
)

var server = gin.Default()

func main() {
	router := server.Group("")
	api.CreateAPIGroup(router)

	server.Run(":8080")
}
