package game

import "fmt"

const (
	SetUp    = 0
	Playing  = 1
	GameOver = 2
)

type GameLogic struct {
	Players   map[string]*PlayerBoard
	GameState int
}

func NewGameLogic() *GameLogic {
	return &GameLogic{
		Players:   make(map[string]*PlayerBoard),
		GameState: SetUp,
	}
}

func (g *GameLogic) AddPlayer(name string) error {
	if g.GameState == SetUp {
		g.Players[name] = NewPlayerBoard(name)

		if len(g.Players) >= 2 && g.GameState == SetUp {
			g.GameState = Playing
			return nil
		}
	}

	return fmt.Errorf("game not set up state")
}

func (g *GameLogic) SetPlayerMine(Id string, x int, y int) (int, error) {
	if g.GameState == SetUp {
		return g.Players[Id].SetMine(x, y)
	}

	return -1, fmt.Errorf("game not set up state")
}

func (g *GameLogic) Shoot(id string, x int, y int) error {
	if g.GameState != Playing {
		for k := range g.Players {
			if k != id {
				return g.Players[id].Shoot(g.Players[k], x, y)
			}
		}
	}

	return fmt.Errorf("game not playing state")
}

func (g *GameLogic) MarkFlag(id string, x int, y int) error {
	if g.GameState != Playing {
		for k := range g.Players {
			if k != id {
				return g.Players[id].MarkFlag(g.Players[k], x, y)
			}
		}
	}

	return fmt.Errorf("game not playing state")
}

func (g *GameLogic) CheckWin() (bool, string) {
	player1 := g.Players[GetKeys(g.Players)[0]]
	player2 := g.Players[GetKeys(g.Players)[1]]

	if player1.Mines <= 0 {
		g.GameState = GameOver
		return true, player2.Id
	} else if player2.Mines <= 0 {
		g.GameState = GameOver
		return true, player1.Id
	}

	return false, ""
}
