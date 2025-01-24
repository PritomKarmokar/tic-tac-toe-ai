package main

import (
	"fmt"
	"math"
)

// Define players
var huPlayer, aiPlayer string

type Move struct {
	Index int
	Score int
}

// Minimax function
func minimax(newBoard []string, player string) Move {

	// Get available spots
	availSpots := emptyIndexes(newBoard)

	// Check terminal states (win, lose, tie)
	if winning(newBoard, huPlayer) {
		return Move{Score: -10}
	} else if winning(newBoard, aiPlayer) {
		return Move{Score: 10}
	} else if len(availSpots) == 0 {
		return Move{Score: 0}
	}

	// Array to collect moves
	var moves []Move

	// Loop through available spots
	for _, spot := range availSpots {
		// Create a move object
		move := Move{Index: spot}

		// Set the empty spot to the current player
		newBoard[spot] = player

		// Recursively call minimax for the opponent
		if player == aiPlayer {
			result := minimax(newBoard, huPlayer)
			move.Score = result.Score
		} else {
			result := minimax(newBoard, aiPlayer)
			move.Score = result.Score
		}

		// Reset the spot to empty
		newBoard[spot] = ""

		// Add the move to the moves array
		moves = append(moves, move)
	}

	// Choose the best move
	var bestMove Move
	if player == aiPlayer {
		bestScore := math.MinInt64
		for _, move := range moves {
			if move.Score > bestScore {
				bestScore = move.Score
				bestMove = move
			}
		}
	} else {
		bestScore := math.MaxInt64
		for _, move := range moves {
			if move.Score < bestScore {
				bestScore = move.Score
				bestMove = move
			}
		}
	}

	return bestMove
}

// Helper functions (to be implemented)

// Returns the indexes of empty spots on the board
func emptyIndexes(board []string) []int {
	var indexes []int
	for i, spot := range board {
		if spot == "" {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// Checks if a player has won
func winning(board []string, player string) bool {
	if (board[0] == player && board[1] == player && board[2] == player) ||
		(board[3] == player && board[4] == player && board[5] == player) ||
		(board[6] == player && board[7] == player && board[8] == player) ||
		(board[0] == player && board[3] == player && board[6] == player) ||
		(board[1] == player && board[4] == player && board[7] == player) ||
		(board[2] == player && board[5] == player && board[8] == player) ||
		(board[0] == player && board[4] == player && board[8] == player) ||
		(board[2] == player && board[4] == player && board[6] == player) {
		return true
	}
	return false
}

func validateMove(board []string, pos int) bool {
	if pos >= 0 && pos < 9 && board[pos] == "" {
		return true
	}
	return false
}

// Print the Tic Tac Toe board
func printBoard(board []string) {
	for i, val := range board {
		if val == "" {
			fmt.Print("_ ")
		} else {
			fmt.Print(val, " ")
		}
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

func main() {

	board := make([]string, 9) // Creating a slice with length of 9

	huPlayer = "O"
	aiPlayer = "X"

	fmt.Println("Initial Game Board")
	printBoard(board)

	for turn := 0; turn < 9; turn++ {
		var pos int
		if turn%2 == 0 {
			for {
				fmt.Print("Your Turn: ")
				_, err := fmt.Scanf("%d", &pos)
				if err != nil {
					fmt.Println("Invalid input, try again.")
				}
				if validateMove(board, pos) {
					board[pos] = huPlayer
					break
				} else {
					fmt.Println("Invalid move, try again.")
				}
			}
		} else {
			var aiTurn Move
			fmt.Print("My Turn: \n")
			aiTurn = minimax(board, aiPlayer)
			board[aiTurn.Index] = aiPlayer
		}

		printBoard(board)

		if winning(board, aiPlayer) {
			fmt.Println("AI wins..")
			break
		} else if winning(board, huPlayer) {
			fmt.Println("You wins..")
			break
		} else if turn == 8 {
			fmt.Println("It's a Tie..")
		}
	}

}
