package api

import (
	client "battleship-server/game"

	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var clients = map[string]*client.Client{}
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func CreateAPIGroup(router *gin.RouterGroup) {
	router.POST("/api/create", createSession)
	router.GET("/api/update/:id", getUpdateSession)
	router.POST("/api/message/:id", sendMessage)
}

func createSession(c *gin.Context) {
	id := randomString(6)
	clients[id] = client.New(id)
	c.String(200, id)
}

func getUpdateSession(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, clients[id])
}

func sendMessage(c *gin.Context) {
	id := c.Param("id")
	msg := c.PostForm("message")
	clients[id].AddMessage(msg)
	c.JSON(200, clients[id])
}
