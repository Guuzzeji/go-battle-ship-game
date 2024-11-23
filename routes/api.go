package routes

import (
	"go-battle-ship-game/game"

	"github.com/gin-gonic/gin"
)

var clients = make(map[string]*game.GameLogic)

type GamePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func CreateAPIGroup(router *gin.RouterGroup) {
	router.POST("/api/create", createGame)
	router.POST("/api/g/:id/join", joinGame)
	router.POST("/api/g/:id/spot/:playerid", gameShoot)
	router.POST("/api/g/:id/add-mine/:playerid", addMine)
	router.POST("/api/g/:id/flag/:playerid", gameFlag)

	router.GET("/api/g/:id/score/:playerid", gameScore)
	router.GET("/api/g/:id/board/:playerid", gameBoard)
	router.GET("/api/g/:id/check", gameCheck)
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

func addMine(c *gin.Context) {
	id := c.Param("id")
	playerId := c.Param("playerid")

	var body GamePosition
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := clients[id].SetPlayerMine(playerId, body.X, body.Y)

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
	clients[id].CheckWin()
	c.JSON(200, clients[id].Players[playerId])
}

func gameFlag(c *gin.Context) {
	id := c.Param("id")
	playerId := c.Param("playerid")

	var body GamePosition
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	clients[id].MarkFlag(playerId, body.X, body.Y)
	clients[id].CheckWin()
	c.JSON(200, clients[id].Players[playerId])
}

func gameScore(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, clients[id])
}

func gameBoard(c *gin.Context) {
	c.JSON(200, clients[c.Param("id")].Players[c.Param("playerid")])
}

func gameCheck(c *gin.Context) {
	_, ok := clients[c.Param("id")]

	if ok {
		c.JSON(200, gin.H{"good": true})
	}

	c.JSON(200, gin.H{"good": false})
}
