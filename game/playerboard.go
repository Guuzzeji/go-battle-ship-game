package game

import (
	"fmt"
)

// Constants for board size and number of mines
const boardSize int = 5
const maxMines int = 5

// Enums for board information
const (
	mine    = -1
	hitMine = -2
	flag    = -3
	show    = -4 // used to show empty spaces (aka zeros, since zero are already being used to show hidden spaces)
)

// Enums for points
const (
	generalPoint  = 1
	hitMinePoint  = -25
	goodFlagPoint = 25
	badFlagPoint  = -15
)

// PlayerBoard represents a player's board
type PlayerBoard struct {
	Id         string
	board      [boardSize][boardSize]int
	InputBoard [boardSize][boardSize]int
	Mines      int
	Points     int
}

// NewPlayerBoard creates and returns a new PlayerBoard with the specified Id.
// The board and InputBoard are initialized to zero values, and the number of Mines
// and Points are set to zero.
func NewPlayerBoard(Id string) *PlayerBoard {
	return &PlayerBoard{
		Id:         Id,
		board:      [boardSize][boardSize]int{},
		InputBoard: [boardSize][boardSize]int{},
		Mines:      0,
		Points:     0,
	}
}

// SetMine places a mine on the board at the specified col and row.  If the
// location is out of bounds or already has a mine, an error is returned.
// If the number of mines on the board has reached the maximum limit, an
// error is returned.  Otherwise, the function returns the number of mines
// remaining that can be placed on the board.
func (pb *PlayerBoard) SetMine(col int, row int) (int, error) {
	if pb.Mines < maxMines {
		if row < 0 || row >= boardSize || col < 0 || col >= boardSize {
			return 0, fmt.Errorf("out of bounds")
		} else if pb.board[row][col] == mine {
			return 0, fmt.Errorf("mine already placed there")
		}

		pb.board[row][col] = mine
		pb.calcBoard()
		pb.Mines++
	} else {
		return 0, fmt.Errorf("max mines reached")
	}

	return maxMines - pb.Mines, nil
}

// calcBoard sets the number of mines adjacent to each cell in the board.
// If a cell has a mine, it is ignored.  Otherwise, the number of mines in
// the adjacent cells is calculated and set in the board.
func (pd *PlayerBoard) calcBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			// Skip mines
			if pd.board[i][j] == mine {
				continue
			}

			numMines := 0 // Sum of adjacent mines from spot

			// Looking up
			if (i+1) < boardSize && pd.board[i+1][j] == mine {
				numMines++
			}
			// Looking down
			if (i-1) >= 0 && pd.board[i-1][j] == mine {
				numMines++
			}
			// Looking right
			if (j+1) < boardSize && pd.board[i][j+1] == mine {
				numMines++
			}
			// Looking left
			if (j-1) >= 0 && pd.board[i][j-1] == mine {
				numMines++
			}

			// Diagonals
			if (i+1 < boardSize) && (j-1 >= 0) && pd.board[i+1][j-1] == mine {
				numMines++
			}
			if (i-1 >= 0) && (j-1 >= 0) && pd.board[i-1][j-1] == mine {
				numMines++
			}
			if (i+1 < boardSize) && (j+1 < boardSize) && pd.board[i+1][j+1] == mine {
				numMines++
			}
			if (i-1 >= 0) && (j+1 < boardSize) && pd.board[i-1][j+1] == mine {
				numMines++
			}

			pd.board[i][j] = numMines
		}
	}
}

// Shoot attempts to hit a spot on the opponent's board at the specified
// column and row. If the location is out of bounds or has already been
// shot, an error is returned. If the spot contains a mine, it is marked
// as hit, and points are deducted. If the spot is empty, it is revealed
// and points are awarded. Adjacent empty spaces are recursively revealed
// as well. If the spot contains a number, it is revealed, and points are
// awarded.
func (pb *PlayerBoard) Shoot(opponent *PlayerBoard, col int, row int) error {
	if row < 0 || row >= boardSize {
		return fmt.Errorf("invalId row")
	} else if col < 0 || col >= boardSize {
		return fmt.Errorf("out of bounds")
	}

	if pb.InputBoard[row][col] != 0 {
		return fmt.Errorf("already shot there")
	}

	if opponent.board[row][col] == mine {
		pb.InputBoard[row][col] = hitMine
		pb.Points += hitMinePoint
		opponent.Mines--
	} else if opponent.board[row][col] == 0 {
		pb.InputBoard[row][col] = -4
		pb.Points += generalPoint

		// Use goroutines to recursively reveal adjacent empty spaces
		go pb.Shoot(opponent, col+1, row)
		go pb.Shoot(opponent, col-1, row)
		go pb.Shoot(opponent, col, row+1)
		go pb.Shoot(opponent, col, row-1)
		go pb.Shoot(opponent, col-1, row-1)
		go pb.Shoot(opponent, col+1, row+1)
		go pb.Shoot(opponent, col-1, row+1)
		go pb.Shoot(opponent, col+1, row-1)

	} else if opponent.board[row][col] > 0 {
		pb.InputBoard[row][col] = opponent.board[row][col]
		pb.Points += generalPoint
	}

	return nil
}

// MarkFlag places a flag on the opponent's board at the specified col and row. If the
// location is out of bounds or has already been shot, an error is returned. If the
// location is a mine, it is marked as hit, and points are awarded. If the location is
// empty, it is marked as a flag, and points are deducted.
func (pb *PlayerBoard) MarkFlag(opponent *PlayerBoard, col int, row int) error {
	if row < 0 || row >= boardSize {
		return fmt.Errorf("invalId row")
	} else if col < 0 || col >= boardSize {
		return fmt.Errorf("out of bounds")
	}

	if pb.InputBoard[row][col] != 0 {
		return fmt.Errorf("already shot there")
	}

	if opponent.board[row][col] == mine {
		pb.InputBoard[row][col] = flag
		pb.Points += goodFlagPoint
		opponent.Mines--
	} else {
		pb.InputBoard[row][col] = flag
		pb.Points += badFlagPoint
	}

	return nil
}
