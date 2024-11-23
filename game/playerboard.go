package game

import (
	"fmt"
)

const boardSize int = 10
const maxMines int = 5

// Enums
const (
	mine    = -1
	hitMine = -2
	flag    = -3
	show    = -4
)

const (
	generalPoint  = 1
	hitMinePoint  = -10
	goodFlagPoint = 25
	badFlagPoint  = -20
)

type PlayerBoard struct {
	Id         string
	Board      [boardSize][boardSize]int
	InputBoard [boardSize][boardSize]int
	Mines      int
	Points     int
}

func NewPlayerBoard(Id string) *PlayerBoard {
	return &PlayerBoard{
		Id:         Id,
		Board:      [boardSize][boardSize]int{},
		InputBoard: [boardSize][boardSize]int{},
		Mines:      0,
		Points:     0,
	}
}

// Function to place ships on the Board
func (pb *PlayerBoard) SetMine(col int, row int) (int, error) {
	if pb.Mines < maxMines {
		if row < 0 || row >= boardSize || col < 0 || col >= boardSize {
			return 0, fmt.Errorf("out of bounds")
		} else if pb.Board[row][col] == mine {
			return 0, fmt.Errorf("mine already placed there")
		}

		pb.Board[row][col] = mine
		pb.calcBoard()
		pb.Mines++
	} else {
		return 0, fmt.Errorf("max mines reached")
	}

	return maxMines - pb.Mines, nil
}

func (pd *PlayerBoard) calcBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if pd.Board[i][j] == mine {
				continue
			}

			numMines := 0
			if (i+1) < boardSize && pd.Board[i+1][j] == mine {
				numMines++
			}

			if (i-1) >= 0 && pd.Board[i-1][j] == mine {
				numMines++
			}

			if (j+1) < boardSize && pd.Board[i][j+1] == mine {
				numMines++
			}

			if (j-1) >= 0 && pd.Board[i][j-1] == mine {
				numMines++
			}

			if (i+1 < boardSize) && (j-1 >= 0) && pd.Board[i+1][j-1] == mine {
				numMines++
			}

			if (i-1 >= 0) && (j-1 >= 0) && pd.Board[i-1][j-1] == mine {
				numMines++
			}

			if (i+1 < boardSize) && (j+1 < boardSize) && pd.Board[i+1][j+1] == mine {
				numMines++
			}

			if (i-1 >= 0) && (j+1 < boardSize) && pd.Board[i-1][j+1] == mine {
				numMines++
			}

			pd.Board[i][j] = numMines
		}
	}
}

// Function to handle shooting
func (pb *PlayerBoard) Shoot(opponent *PlayerBoard, col int, row int) error {
	if row < 0 || row >= boardSize {
		return fmt.Errorf("invalId row")
	} else if col < 0 || col >= boardSize {
		return fmt.Errorf("out of bounds")
	}

	if pb.InputBoard[row][col] != 0 {
		return fmt.Errorf("already shot there")
	}

	if opponent.Board[row][col] == mine {
		pb.InputBoard[row][col] = hitMine
		pb.Points += hitMinePoint
		opponent.Mines--
	} else if opponent.Board[row][col] == 0 {
		pb.InputBoard[row][col] = -4
		pb.Points += generalPoint

		go pb.Shoot(opponent, col+1, row)
		go pb.Shoot(opponent, col-1, row)
		go pb.Shoot(opponent, col, row+1)
		go pb.Shoot(opponent, col, row-1)
		go pb.Shoot(opponent, col-1, row-1)
		go pb.Shoot(opponent, col+1, row+1)
		go pb.Shoot(opponent, col-1, row+1)
		go pb.Shoot(opponent, col+1, row-1)

	} else if opponent.Board[row][col] > 0 {
		pb.InputBoard[row][col] = opponent.Board[row][col]
		pb.Points += goodFlagPoint
	}

	return nil
}

func (pb *PlayerBoard) MarkFlag(opponent *PlayerBoard, col int, row int) error {
	if row < 0 || row >= boardSize {
		return fmt.Errorf("invalId row")
	} else if col < 0 || col >= boardSize {
		return fmt.Errorf("out of bounds")
	}

	if pb.InputBoard[row][col] != 0 {
		return fmt.Errorf("already shot there")
	}

	if opponent.Board[row][col] == mine {
		pb.InputBoard[row][col] = flag
		pb.Points += goodFlagPoint
		opponent.Mines--
	} else {
		pb.InputBoard[row][col] = flag
		pb.Points += badFlagPoint
	}

	return nil
}
