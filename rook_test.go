package main

import (
	"testing"
)

func TestRookGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	rook1, _ := white.getPieceByCoordinate(1, 1)

	// Initially, the white rook can't move because it's boxed in.
	moves := rook1.generateValidMoves(Square{x: 1, y: 1})
	if len(moves) > 0 {
		t.Errorf("Rook in starting position has valid moves.")
	}

	// Sneakiness to force a move without validation.
	genMoves := rook1.generateValidMoves(Square{x: 4, y: 4})

	// Rook can move vertically up to the first opponent piece.
	// Rook can move horizontally across the board.
	expMoves := []*Square{
		&Square{x: 4, y: 3},
		&Square{x: 4, y: 5},
		&Square{x: 4, y: 6},
		&Square{x: 4, y: 7},
		&Square{x: 1, y: 3},
		&Square{x: 2, y: 3},
		&Square{x: 3, y: 3},
		&Square{x: 5, y: 3},
		&Square{x: 6, y: 3},
		&Square{x: 7, y: 3},
		&Square{x: 8, y: 3},
	}

	arraysMatch, err := squareArraysMatch(genMoves, expMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}
