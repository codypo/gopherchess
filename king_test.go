package main

import (
	"testing"
)

func TestKingGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	king1 := b.getPieceByCoordinates(5, 1)

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

func TestKingCanCapture(t *testing.T) {
	b := NewBoard()
	wKing := b.getPieceByCoordinates(5, 1)
	bPawn := b.getPieceByCoordinates(7, 7)

	bPawn.move(&Square{x: 7, y: 5})
	if bPawn.captured {
		t.Errorf("Pawn started out in captured state.")
	}

	// Should not be able to capture pieces it cannot access.
	moveErr := wKing.move(&Square{x: 7, y: 5})
	if moveErr == nil {
		t.Errorf("Invalid move allowed for capture.")
	}

	wKing.forceMove(&Square{x: 6, y: 4})
	wKing.move(&Square{x: 7, y: 5})
	if !bPawn.captured {
		t.Errorf("Captured pawn not in captured state.")
	}

	if bPawn.getSquare() != nil {
		t.Errorf("Captured pawn has a non-nil position.")
	}
}
