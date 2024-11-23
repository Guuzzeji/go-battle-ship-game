package routes

import (
	"go-battle-ship-game/game"

	"github.com/gin-gonic/gin"
)

var clients = make(map[string]*game.GameLogic)

type GamePosition struct {
	X int    `json:"x"`
	Y string `json:"y"`
}

func CreateAPIGroup(router *gin.RouterGroup) {
	router.POST("/api/create-game", createGame)
	router.POST("/api/g/:id/join-game", joinGame)
	router.POST("/api/g/:id/shoot/:playerid", gameShoot)
	router.POST("/api/g/:id/add-ship/:playerid", addShip)
	router.GET("/api/g/:id/board/:playerid", getBoard)
}

func createGame(c *gin.Context) {
	id := RandomString(6)
	clients[id] = game.NewGameLogic()
	c.JSON(200, gin.H{"id": id})
}

func joinGame(c *gin.Context) {
	gameId := c.Param("id")
	id := RandomString(4)
	clients[gameId].AddPlayer(id)
	c.JSON(200, gin.H{"id": id})
}

func getBoard(c *gin.Context) {
	id := c.Param("id")
	playerId := c.Param("playerid")
	c.JSON(200, clients[id].Players[playerId])
}

func addShip(c *gin.Context) {
	id := c.Param("id")
	playerId := c.Param("playerid")

	var body GamePosition
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := clients[id].SetShips(playerId, body.X, body.Y)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, clients[id].Players[playerId])
}

func gameShoot(c *gin.Context) {
	id := c.Param("id")
	playerId := c.Param("playerid")

	var body GamePosition
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	clients[id].Shoot(playerId, body.X, body.Y)
	c.JSON(200, clients[id])
}
