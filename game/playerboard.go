package game

import (
	"fmt"
)

const BoardSize int = 10
const maxShips int = 3

// Enums
const (
	shipSymbol = 1
	hitSymbol  = 2
	missSymbol = -1
)

type PlayerBoard struct {
	Id         string
	Board      [BoardSize][BoardSize]int
	InputBoard [BoardSize][BoardSize]int
	ShipPlaced int
	Points     int
}

func NewPlayerBoard(Id string) *PlayerBoard {
	return &PlayerBoard{
		Id:         Id,
		Board:      [BoardSize][BoardSize]int{},
		InputBoard: [BoardSize][BoardSize]int{},
		ShipPlaced: 0,
		Points:     0,
	}
}

// Function to place ships on the Board
func (pb *PlayerBoard) SetShips(x int, y string) (bool, error) {
	var row, col int
	col = x

	if pb.ShipPlaced < maxShips {
		row = ConvertLetterToRow(y)

		fmt.Println(row, col)

		if row < 0 || row >= BoardSize || col < 0 || col >= BoardSize {
			return false, fmt.Errorf("out of bounds")
		} else if pb.Board[row][col] == shipSymbol {
			return false, fmt.Errorf("ship already placed there")
		}

		pb.Board[row][col] = shipSymbol
		pb.ShipPlaced += 1
	}

	return true, nil
}

// Function to handle shooting
func (pb *PlayerBoard) Shoot(opponent *PlayerBoard, x int, y string) (int, error) {
	var row, col int
	col = x

	row = ConvertLetterToRow(y)
	if row < 0 {
		return -1, fmt.Errorf("invalId row")
	}

	if col < 0 || col >= BoardSize {
		return -1, fmt.Errorf("out of bounds")
	}

	if opponent.Board[row][col] != 0 {
		return -1, fmt.Errorf("already shot there")
	}

	// If valId coordinates, process the shot
	if opponent.Board[row][col] == shipSymbol {
		opponent.InputBoard[row][col] = hitSymbol
		opponent.Board[row][col] = 0
		opponent.ShipPlaced--
		return 1, nil
	}

	opponent.InputBoard[row][col] = missSymbol
	return 0, nil
}
