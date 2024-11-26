package game

import "fmt"

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

func NewGameLogic() *GameLogic {
	return &GameLogic{
		PlayerOne:        NewPlayerBoard("1"),
		PlayerTwo:        NewPlayerBoard("2"),
		isPlayerOneReady: false,
		isPlayerTwoReady: false,
		GameState:        setup,
	}
}

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

func (g *GameLogic) CheckWin() (bool, string) {
	if g.PlayerOne.Mines <= 0 && g.PlayerTwo.Mines <= 0 {
		if g.PlayerOne.Points > g.PlayerTwo.Points {
			g.GameState = gameover
			return true, "1"

		} else if g.PlayerOne.Points < g.PlayerTwo.Points {
			g.GameState = gameover
			return true, "2"
		}
	}

	return false, ""
}
