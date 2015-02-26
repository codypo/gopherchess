package main

import (
	"testing"
)

func TestKingGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	king1, _ := white.getPieceByCoordinate(5, 1)

	// Initially, the white king can't move because it's boxed in.
	moves := king1.generateMoves(Square{x: 5, y: 1})
	if len(moves) > 0 {
		t.Errorf("King in starting position has valid moves.")
	}

	// Sneakiness to force a move without validation.
	moves = king1.generateMoves(Square{x: 4, y: 4})

	// We selected a starting square so we can move in all directions.
	expMoves := []*Square{
		&Square{x: 5, y: 4},
		&Square{x: 3, y: 4},
		&Square{x: 4, y: 5},
		&Square{x: 4, y: 3},
		&Square{x: 5, y: 3},
		&Square{x: 5, y: 5},
		&Square{x: 3, y: 5},
		&Square{x: 3, y: 3},
	}

	arraysMatch, err := squaresMatch(moves, expMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}
