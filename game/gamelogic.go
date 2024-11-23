package game

import "fmt"

const (
	SetUp    = 0
	Playing  = 1
	GameOver = 2
)

type GameLogic struct {
	Players     map[string]*PlayerBoard
	GameState   int
	currentTurn string
	opponent    string
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

			m := GetKeys(g.Players)

			g.currentTurn = g.Players[m[0]].Id
			g.opponent = g.Players[m[1]].Id
			return nil
		}
	}

	return fmt.Errorf("game not set up state")
}

func (g *GameLogic) SetShips(Id string, x int, y string) (bool, error) {
	if g.GameState == SetUp {
		return g.Players[Id].SetShips(x, y)
	}

	return false, fmt.Errorf("game not set up state")
}

func (g *GameLogic) Shoot(Id string, x int, y string) (int, error) {
	if g.GameState != Playing {
		for k := range g.Players {
			if k != Id {
				return g.Players[k].Shoot(g.Players[k], x, y)
			}
		}
	}

	return -1, fmt.Errorf("game not playing state")
}

func (g *GameLogic) CheckWin() (bool, string) {
	player1 := g.Players[GetKeys(g.Players)[0]]
	player2 := g.Players[GetKeys(g.Players)[1]]

	if player1.ShipPlaced <= 0 {
		g.GameState = GameOver
		return true, player2.Id
	} else if player2.ShipPlaced <= 0 {
		g.GameState = GameOver
		return true, player1.Id
	}

	return false, ""
}
