package main

import (
	"testing"
)

func TestBishopGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	bishop1, _ := white.getPieceByCoordinate(3, 1)

	// Initially, the white bishop can't move because it's boxed in.
	moves := bishop1.generateValidMoves(Square{x: 3, y: 1})
	if len(moves) > 0 {
		t.Errorf("Bishop in starting position has valid moves.")
	}

	// Sneakiness to force a move without validation.
	genMoves := bishop1.generateValidMoves(Square{x: 3, y: 4})

	// We selected a starting square so we can move up left, up right, down left, down right.
	expMoves := []*Square{
		&Square{x: 2, y: 5},
		&Square{x: 1, y: 6},
		&Square{x: 4, y: 5},
		&Square{x: 5, y: 6},
		&Square{x: 2, y: 3},
		&Square{x: 4, y: 3},
		&Square{x: 6, y: 7},
	}

	arraysMatch, err := squareArraysMatch(genMoves, expMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}
