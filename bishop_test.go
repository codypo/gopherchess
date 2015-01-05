package main

import (
	"fmt"
	"testing"
)

func TestBishopGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	bishop1 := white.getPieceByCoordinate(3, 1)

	// Initially, the white bishop can't move because it's boxed in.
	moves := bishop1.generateValidMoves(Square{x: 3, y: 1})
	if len(moves) > 0 {
		t.Errorf("Bishop in starting position has valid moves.")
	}

	fmt.Println("Done checking the starting position.")

	// Sneakiness to force a move without validation.
	genMoves := bishop1.generateValidMoves(Square{x: 3, y: 4})

	var expMoves [7]*Square
	// We selected a starting square so we can move up left, up right, down left, down right.
	expMoves[0] = &Square{x: 2, y: 5}
	expMoves[1] = &Square{x: 1, y: 6}
	expMoves[2] = &Square{x: 4, y: 5}
	expMoves[3] = &Square{x: 5, y: 6}
	expMoves[4] = &Square{x: 2, y: 3}
	expMoves[5] = &Square{x: 4, y: 3}
	expMoves[6] = &Square{x: 6, y: 7}

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
