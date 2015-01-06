package main

import (
	"fmt"
	"testing"
)

func TestRookGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	rook1 := white.getPieceByCoordinate(1, 1)

	// Initially, the white rook can't move because it's boxed in.
	moves := rook1.generateValidMoves(Square{x: 1, y: 1})
	if len(moves) > 0 {
		t.Errorf("Rook in starting position has valid moves.")
	}

	fmt.Println("Done checking the starting position.")

	// Sneakiness to force a move without validation.
	genMoves := rook1.generateValidMoves(Square{x: 1, y: 3})

	var expMoves [11]*Square
	// Rook can move horizontally across the board.
	expMoves[0] = &Square{x: 2, y: 3}
	expMoves[1] = &Square{x: 3, y: 3}
	expMoves[2] = &Square{x: 4, y: 3}
	expMoves[3] = &Square{x: 5, y: 3}
	expMoves[4] = &Square{x: 6, y: 3}
	expMoves[5] = &Square{x: 7, y: 3}
	expMoves[6] = &Square{x: 8, y: 3}

	// Rook can move vertically up to the first opponent piece.
	expMoves[7] = &Square{x: 1, y: 4}
	expMoves[8] = &Square{x: 1, y: 5}
	expMoves[9] = &Square{x: 1, y: 6}
	expMoves[10] = &Square{x: 1, y: 7}

	if len(genMoves) != len(expMoves) {
		t.Errorf("Generated moves are of length %d, while expected moves are of length %d", len(genMoves), len(expMoves))
	}

	// Certainly a better way to do this, but here we are.
	for _, expMove := range expMoves {
		foundMatch := false
		for _, genMove := range genMoves {
			if expMove.equals(*genMove) {
				foundMatch = true
				break
			}
		}

		if !foundMatch {
			t.Errorf("Expected to find %d, %d in list of generated moves, but did not.", expMove.x, expMove.y)
		}
	}

}
