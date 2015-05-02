package main

import "fmt"

// Main entry point to the game.
func main() {
	board := setupGame()
	gameLoop(board)
}

// Instantiates the game and collects player info.
func setupGame() *Board {

	fmt.Printf("Hello, puny human.  What is your name?  ")
	var playerName string
	for {
		_, err := fmt.Scanf("%s", &playerName)
		if err == nil && len(playerName) > 0 {
			break
		}
	}

	fmt.Printf("The game is afoot, %s.  You are white and move first.\n", playerName)

	b := NewBoard()
	b.setPlayerName(White, playerName)
	b.setPlayerName(Black, "gopherchess")
	return b
}

// The loop that's processing user input, applying it to the
// game, and receiving moves from the gopherchess algorithms.
func gameLoop(b *Board) {
	// TODO: Let's color this via http://godoc.org/github.com/daviddengcn/go-colortext
	b.prettyPrint()

	for {
		switch b.getGameState() {
		case GameOn: // Play on, player.
			processMove(b)
			break
		case WhiteInCheck: // Uh oh, get out of check, white.
			fmt.Printf("White is checked!\n")
			processMove(b)
			break
		case BlackInCheck: // Uh oh, you get out of check right now, black.
			fmt.Printf("Black is checked!\n")
			processMove(b)
			break
		case WhiteCheckmated: // Game over; black wins.
			fmt.Printf("Black has checkmated white!  Black wins.  Congrats, you get 11 gopherchess bucks, redeemable for nothing.\n")
			return
		case BlackCheckmated: // Game over; white wins.
			fmt.Printf("White has checkmated black!  White wins.  Congrats, you get 5 gopherchess bucks, redeemable for nothing.\n")
			return
		case Draw: // Game over; great sadness.
			fmt.Printf("DRAW.  No one wins, much like life.\n")
			return
		}
	}
}

// Validates and applies a user move.
func processMove(b *Board) {
	fmt.Printf("Your move, %s. [Example: move bishop to c3 with Bc3]\n", b.getPlayerWithNextMove())
	var move string
	for {
		_, err := fmt.Scanf("%s", &move)
		if err == nil && len(move) > 0 {
			break
		}

	}

	// TODO: Now, do something!
	fmt.Printf("I should do something here, but I don't yet.\n")
}
