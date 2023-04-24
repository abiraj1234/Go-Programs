package main

import (
	"fmt"
)

func PrintXoGame(Board [3][3]string) {
	fmt.Println(" | 1 2 3")
	fmt.Println("-+-------")
	for i := 0; i < len(Board); i++ {
		fmt.Print(i+1, "| ")
		for j := 0; j < len(Board); j++ {
			fmt.Print(Board[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	// Initialize the game board
	XoGameBoard := [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	player := "X"
	for {
		PrintXoGame(XoGameBoard)
		// Prompt the player to enter their move
		var row, col int

		fmt.Printf("Player %s, enter your move (row column): ", player)

		_, input := fmt.Scanf("%d %d\n", &row, &col)
		if input != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		// Check if the move is valid
		if row < 1 || row > 3 || col < 1 || col > 3 || XoGameBoard[row-1][col-1] != " " {
			fmt.Println("Invalid move. Try again.")
			continue
		}

		// Make the move and check for a win
		XoGameBoard[row-1][col-1] = player
		if ChekWinPlayer(XoGameBoard, player) {
			PrintXoGame(XoGameBoard)
			fmt.Printf("Player %s wins!\n", player)
			break
		}

		// Check for a tie game
		if IfXoBoardIsFull(XoGameBoard) {
			fmt.Println("Tie game.")
			PrintXoGame(XoGameBoard)
			break
		}

		// Switch to the other player
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
}

func ChekWinPlayer(board [3][3]string, player string) bool {
	for i := 0; i < 3; i++ {
		switch {
		//chek row
		case board[i][0] == player && board[i][1] == player && board[i][2] == player:
			return true
		// Check for three in a column
		case board[0][i] == player && board[1][i] == player && board[2][i] == player:
			return true
		}
	}

	// Check for three on a diagonal
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	// No win found
	return false
}

// Check if the board is full
func IfXoBoardIsFull(board [3][3]string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}
