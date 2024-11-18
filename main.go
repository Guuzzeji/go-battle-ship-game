package main

import (
	"fmt"
	"strings"
)

const boardSize int = 10
const shipSymbol = 1  // Symbol for a ship
const hitSymbol = 2   // Symbol for a hit
const missSymbol = -1 // Symbol for a miss

type PlayerBoard struct {
	name          string
	board         [boardSize][boardSize]int
	shootingBoard [boardSize][boardSize]int
	shipsLeft     int
}

// Function to display the player's board
func (pb *PlayerBoard) displayBoard() {
	fmt.Println(pb.name + "'s Board")
	fmt.Print("  ")
	for i := 0; i < boardSize; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 0; i < boardSize; i++ {
		fmt.Printf("%c ", 'A'+rune(i))
		for j := 0; j < boardSize; j++ {
			if pb.board[i][j] == 0 {
				fmt.Print(". ") // Empty space
			} else {
				fmt.Printf("%d ", pb.board[i][j]) // Ship or other symbols
			}
		}
		fmt.Println()
	}
}

// Function to display the shooting board
func (pb *PlayerBoard) displayShootingBoard() {
	fmt.Println(pb.name + "'s Shooting Board")
	fmt.Print("  ")
	for i := 0; i < boardSize; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 0; i < boardSize; i++ {
		fmt.Printf("%c ", 'A'+rune(i))
		for j := 0; j < boardSize; j++ {
			switch pb.shootingBoard[i][j] {
			case 0:
				fmt.Print(". ") // Empty space
			case hitSymbol:
				fmt.Print("X ") // Hit
			case missSymbol:
				fmt.Print("O ") // Miss
			}
		}
		fmt.Println()
	}
}

// Function to place ships on the board
func (pb *PlayerBoard) setShips() {
	numShipsToPlace := 3
	var cords string
	pb.shipsLeft = numShipsToPlace
	var row int
	var col int

	fmt.Printf("\n%s, place your ships!\n", pb.name)
	pb.displayBoard()

	for numShipsToPlace > 0 {
		fmt.Printf("\nNumber of ships to be placed: %d\n", numShipsToPlace)
		fmt.Println("Please enter the coordinates for each ship (e.g., D4):")
		fmt.Scan(&cords)

		cords = strings.ToUpper(cords)

		if len(cords) != 2 {
			fmt.Println("Invalid coordinates, please enter in format (e.g., D4).")
			continue
		}

		switch cords[0] {
		case 'A':
			row = 0
		case 'B':
			row = 1
		case 'C':
			row = 2
		case 'D':
			row = 3
		case 'E':
			row = 4
		case 'F':
			row = 5
		case 'G':
			row = 6
		case 'H':
			row = 7
		case 'I':
			row = 8
		case 'J':
			row = 9
		default:
			fmt.Println("Invalid row letter. Please enter a letter from A-J.")
			continue
		}

		col = int(cords[1] - '0')

		if row < 0 || row >= boardSize || col < 0 || col >= boardSize {
			fmt.Println("Invalid coordinates. Please enter a letter from A-J and a number from 0-9.")
		} else if pb.board[row][col] == shipSymbol {
			fmt.Println("Ship already placed there. Try again.")
		} else {
			pb.board[row][col] = shipSymbol
			numShipsToPlace--
		}
	}

	pb.displayBoard()
}

// Function to handle shooting
func (pb *PlayerBoard) shoot(opponent *PlayerBoard) {
	var cords string
	var row, col int
	var validShot bool

	for !validShot { // Repeat until a valid shot is made
		fmt.Printf("\n%s's turn to shoot!\n", pb.name)
		opponent.displayShootingBoard()
		fmt.Println("Please enter the coordinates to shoot (e.g., D4):")
		fmt.Scan(&cords)

		cords = strings.ToUpper(cords)

		if len(cords) != 2 {
			fmt.Println("Invalid input format. Enter coordinates like D4.")
			continue // Prompt again
		}

		// Convert letter and number to row and col
		switch cords[0] {
		case 'A':
			row = 0
		case 'B':
			row = 1
		case 'C':
			row = 2
		case 'D':
			row = 3
		case 'E':
			row = 4
		case 'F':
			row = 5
		case 'G':
			row = 6
		case 'H':
			row = 7
		case 'I':
			row = 8
		case 'J':
			row = 9
		default:
			fmt.Println("Invalid row letter. Enter a letter from A-J.")
			continue // Prompt again
		}

		col = int(cords[1] - '0')

		if col < 0 || col >= boardSize {
			fmt.Println("Invalid column number. Enter a number from 0-9.")
			continue // Prompt again
		}

		if opponent.shootingBoard[row][col] != 0 {
			fmt.Println("You've already shot here. Try another spot.")
			continue // Prompt again
		}

		// If valid coordinates, process the shot
		validShot = true // Exit the loop after this shot
		if opponent.board[row][col] == shipSymbol {
			fmt.Println("Hit!")
			opponent.shootingBoard[row][col] = hitSymbol
			opponent.board[row][col] = 0
			opponent.shipsLeft--
		} else {
			fmt.Println("Miss!")
			opponent.shootingBoard[row][col] = missSymbol
		}
	}

}


func main() {
	var player1Name, player2Name string

	fmt.Print("Welcome! Player 1, enter your name: ")
	fmt.Scan(&player1Name)
	fmt.Print("Player 2, enter your name: ")
	fmt.Scan(&player2Name)

	player1 := PlayerBoard{name: player1Name}
	player2 := PlayerBoard{name: player2Name}

	// Both players place their ships
	player1.setShips()
	player2.setShips()

	currentTurn := &player1
	opponent := &player2

	// Game loop
	for player1.shipsLeft > 0 && player2.shipsLeft > 0 {
		currentTurn.shoot(opponent)
		// Switch turns
		if currentTurn == &player1 {
			currentTurn = &player2
			opponent = &player1
		} else {
			currentTurn = &player1
			opponent = &player2
		}
	}

	if player1.shipsLeft == 0 {
		fmt.Printf("\n%s wins!\n", player2.name)
	} else {
		fmt.Printf("\n%s wins!\n", player1.name)
	}
}
