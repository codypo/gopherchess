package main

import (
	"testing"
)

func TestQueenGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	queen1 := white.getPieceByCoordinate(4, 1)

	// Initially, the white queen can't move because it's boxed in.
	moves := queen1.generateValidMoves(Square{x: 1, y: 1})
	if len(moves) > 0 {
		t.Errorf("Queen in starting position has valid moves.")
	}

	// Sneakiness to force a move without validation.
	genMoves := queen1.generateValidMoves(Square{x: 4, y: 4})

	var expMoves [19]*Square

	// Queen can move vertically up to the first opponent piece.
	// Queen can move horizontally across the board.
	expMoves[0] = &Square{x: 4, y: 3}
	expMoves[1] = &Square{x: 4, y: 5}
	expMoves[2] = &Square{x: 4, y: 6}
	expMoves[3] = &Square{x: 4, y: 7}
	expMoves[4] = &Square{x: 1, y: 3}
	expMoves[5] = &Square{x: 2, y: 3}
	expMoves[6] = &Square{x: 3, y: 3}
	expMoves[7] = &Square{x: 5, y: 3}
	expMoves[8] = &Square{x: 6, y: 3}
	expMoves[9] = &Square{x: 7, y: 3}
	expMoves[10] = &Square{x: 8, y: 3}
	expMoves[11] = &Square{x: 5, y: 5}
	expMoves[12] = &Square{x: 6, y: 6}
	expMoves[13] = &Square{x: 5, y: 3}
	expMoves[14] = &Square{x: 6, y: 2}
	expMoves[15] = &Square{x: 3, y: 3}
	expMoves[16] = &Square{x: 5, y: 3}
	expMoves[17] = &Square{x: 7, y: 1}
	expMoves[18] = &Square{x: 7, y: 7}

	arraysMatch, err := squareArraysMatch(genMoves, expMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}
