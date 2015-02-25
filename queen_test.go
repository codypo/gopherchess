package main

import (
	"testing"
)

func TestQueenGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	queen1, _ := white.getPieceByCoordinate(4, 1)

	// Initially, the white queen can't move because it's boxed in.
	moves := queen1.generateValidMoves(Square{x: 1, y: 1})
	if len(moves) > 0 {
		t.Errorf("Queen in starting position has valid moves.")
	}

	// Sneakiness to force a move without validation.
	genMoves := queen1.generateValidMoves(Square{x: 4, y: 4})

	// Queen can move vertically, horizontally, and diagonally.  So regal!
	expMoves := []*Square{
		&Square{x: 4, y: 3},
		&Square{x: 4, y: 5},
		&Square{x: 4, y: 6},
		&Square{x: 4, y: 7},
		&Square{x: 1, y: 4},
		&Square{x: 2, y: 4},
		&Square{x: 3, y: 4},
		&Square{x: 5, y: 4},
		&Square{x: 6, y: 4},
		&Square{x: 7, y: 4},
		&Square{x: 8, y: 4},
		&Square{x: 5, y: 5},
		&Square{x: 6, y: 6},
		&Square{x: 5, y: 3},
		&Square{x: 2, y: 6},
		&Square{x: 3, y: 3},
		&Square{x: 3, y: 5},
		&Square{x: 1, y: 7},
		&Square{x: 7, y: 7},
	}

	arraysMatch, err := squareArraysMatch(genMoves, expMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}
