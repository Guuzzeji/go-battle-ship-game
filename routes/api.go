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

// CreateAPIGroup adds all the main routes to the API router.
//
// It adds the following endpoints:
//
//   - POST /api/create: Creates a new game session and returns the session ID.
//
//   - POST /api/g/:id/join: Allows a player to join a game session.
//
//   - POST /api/g/:id/spot/:playerid: Places a move on the board.
//
//   - POST /api/g/:id/add-mine/:playerid: Adds a mine to the board during setup.
//
//   - POST /api/g/:id/flag/:playerid: Places a flag on the board.
//
//   - GET /api/g/:id/score/:playerid: Returns the score of the game.
//
//   - GET /api/g/:id/board/:playerid: Returns the game board.
//
//   - GET /api/g/:id/check: Checks if the game session is valid.
func CreateAPIGroup(router *gin.RouterGroup) {
	router.POST("/api/create", createGameSession)
	router.POST("/api/g/:id/join", joinGameSession)
	router.POST("/api/g/:id/spot/:playerid", placeSpot)
	router.POST("/api/g/:id/add-mine/:playerid", addMine)
	router.POST("/api/g/:id/flag/:playerid", placeFlag)

	router.GET("/api/g/:id/score/:playerid", getScore)
	router.GET("/api/g/:id/board/:playerid", getPlayerBoard)
	router.GET("/api/g/:id/check", checkSession)
}

func createGameSession(c *gin.Context) {
	sessionId := randomString(6)
	clients[sessionId] = game.NewGameLogic()
	c.JSON(200, gin.H{"id": sessionId})
}

func joinGameSession(c *gin.Context) {
	sessionId := c.Param("id")
	if _, ok := clients[sessionId]; !ok {
		c.JSON(400, gin.H{"error": "game not found"})
		return
	}

	playerId, err := clients[sessionId].AddPlayer()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"id": playerId})
}

func addMine(c *gin.Context) {
	sessionId := c.Param("id")
	playerId := c.Param("playerid")

	if _, ok := clients[sessionId]; !ok {
		c.JSON(400, gin.H{"error": "game not found"})
		return
	}

	var body GamePosition
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := clients[sessionId].SetPlayerMine(playerId, body.X, body.Y)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"done": true})
}

func placeSpot(c *gin.Context) {
	sessionId := c.Param("id")
	playerId := c.Param("playerid")

	if _, ok := clients[sessionId]; !ok {
		c.JSON(400, gin.H{"error": "game not found"})
		return
	}

	var body GamePosition
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := clients[sessionId].Shoot(playerId, body.X, body.Y)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	clients[sessionId].CheckWin()
	c.JSON(200, gin.H{"done": true})
}

func placeFlag(c *gin.Context) {
	sessionId := c.Param("id")
	playerId := c.Param("playerid")

	if _, ok := clients[sessionId]; !ok {
		c.JSON(400, gin.H{"error": "game not found"})
		return
	}

	var body GamePosition
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := clients[sessionId].MarkFlag(playerId, body.X, body.Y)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	clients[sessionId].CheckWin()
	c.JSON(200, gin.H{"done": true})
}

func getScore(c *gin.Context) {
	sessionId := c.Param("id")
	if _, ok := clients[sessionId]; !ok {
		c.JSON(400, gin.H{"error": "game not found"})
		return
	}

	c.JSON(200, gin.H{
		"player1":      clients[sessionId].PlayerOne.Points,
		"player2":      clients[sessionId].PlayerTwo.Points,
		"gameState":    clients[sessionId].GameState,
		"player1mines": clients[sessionId].PlayerOne.Mines,
		"player2mines": clients[sessionId].PlayerTwo.Mines,
	})
}

func getPlayerBoard(c *gin.Context) {
	sessionId := c.Param("id")
	playerId := c.Param("playerid")

	if _, ok := clients[sessionId]; !ok {
		c.JSON(400, gin.H{"error": "game not found"})
		return
	}

	if playerId != "2" && playerId != "1" {
		c.JSON(400, gin.H{"error": "player id not found"})
		return
	}

	if playerId == "1" {
		c.JSON(200, gin.H{"board": clients[sessionId].PlayerOne.InputBoard})
		return
	}

	c.JSON(200, gin.H{"board": clients[sessionId].PlayerTwo.InputBoard})

}

func checkSession(c *gin.Context) {
	_, ok := clients[c.Param("id")]
	if ok {
		c.JSON(200, gin.H{"good": true})
	}
	c.JSON(400, gin.H{"good": false})
}
