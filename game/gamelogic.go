package game

import "fmt"

// Enums for game state
const (
	setup    = 0
	playing  = 1
	gameover = 2
)

type GameLogic struct {
	PlayerOne        *PlayerBoard
	PlayerTwo        *PlayerBoard
	isPlayerOneReady bool
	isPlayerTwoReady bool
	GameState        int
}

// NewGameLogic initializes and returns a new GameLogic instance.
// It sets up two player boards, PlayerOne and PlayerTwo, each with their initial settings.
// The game state is initialized to the setup phase, and both players are marked as not ready.
func NewGameLogic() *GameLogic {
	return &GameLogic{
		PlayerOne:        NewPlayerBoard("1"),
		PlayerTwo:        NewPlayerBoard("2"),
		isPlayerOneReady: false,
		isPlayerTwoReady: false,
		GameState:        setup,
	}
}

// AddPlayer adds a player to the game.
// If the game is in the setup state, it adds the player and returns the player's id.
// If the game is not in the setup state, it returns an error.
// If both players are already present, it returns an error.
func (g *GameLogic) AddPlayer() (string, error) {
	if g.GameState == setup {
		if !g.isPlayerOneReady {
			g.isPlayerOneReady = true
			return "1", nil
		} else if !g.isPlayerTwoReady {
			g.isPlayerTwoReady = true
			return "2", nil
		}
	}

	return "", fmt.Errorf("game not set up state")
}

// SetPlayerMine adds a mine to the specified player's board.
// If the game is in the setup state, it adds the mine and checks if the game can be started.
// If both players have reached the maximum number of mines, the game is advanced to the
// playing state.
// If the game is not in the setup state, it returns an error.
func (g *GameLogic) SetPlayerMine(id string, x int, y int) error {
	if g.GameState == setup {
		if id[0] == '1' {
			_, err := g.PlayerOne.SetMine(x, y)

			if g.PlayerOne.Mines >= maxMines && g.PlayerTwo.Mines >= maxMines {
				g.GameState = playing
			}

			return err
		} else {
			_, err := g.PlayerTwo.SetMine(x, y)

			if g.PlayerOne.Mines >= maxMines && g.PlayerTwo.Mines >= maxMines {
				g.GameState = playing
			}

			return err
		}
	}

	return fmt.Errorf("game not set up state")
}

// Shoot is used by a player to shoot a location on their opponent's board.
// If the game is in the playing state, it calls the Shoot function on the
// corresponding player's board. If the game is not in the playing state,
// it returns an error.
func (g *GameLogic) Shoot(id string, x int, y int) error {
	if g.GameState == playing {
		if id == "1" {
			return g.PlayerOne.Shoot(g.PlayerTwo, x, y)
		} else {
			return g.PlayerTwo.Shoot(g.PlayerOne, x, y)
		}
	}

	return fmt.Errorf("game not playing state")
}

// MarkFlag is used by a player to mark a location on their opponent's board as a mine.
// If the game is in the playing state, it calls the MarkFlag function on the
// corresponding player's board. If the game is not in the playing state,
// it returns an error.
func (g *GameLogic) MarkFlag(id string, x int, y int) error {
	if g.GameState == playing {
		if id == "1" {
			return g.PlayerOne.MarkFlag(g.PlayerTwo, x, y)
		} else {
			return g.PlayerTwo.MarkFlag(g.PlayerOne, x, y)
		}
	}

	return fmt.Errorf("game not playing state")
}

// CheckWin checks if the game is over and returns true if it is, and the
// id of the player who won (either "1" or "2"). If the game is not over, it
// returns false and an empty string.
func (g *GameLogic) CheckWin() (bool, string) {
	if g.PlayerOne.Mines <= 0 && g.PlayerTwo.Mines <= 0 {
		g.GameState = gameover
		if g.PlayerOne.Points > g.PlayerTwo.Points {
			return true, "1"
		} else if g.PlayerOne.Points < g.PlayerTwo.Points {
			return true, "2"
		}
		return true, "0"
	}
	return false, ""
}
