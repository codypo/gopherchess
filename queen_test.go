package main

import (
	"testing"
)

func TestQueenGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	queen1, _ := white.getPieceByCoordinate(4, 1)

	// Initially, the white queen can't move because it's boxed in.
	moves := queen1.generateMoves(Square{x: 1, y: 1})
	if len(moves) > 0 {
		t.Errorf("Queen in starting position has valid moves.")
	}

	// Sneakiness to force a move without validation.
	genMoves := queen1.generateMoves(Square{x: 4, y: 4})

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

	arraysMatch, err := squaresMatch(genMoves, expMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}

func TestQueenCanCapture(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)
	wQueen, _ := white.getPieceByCoordinate(4, 1)
	bPawn, _ := black.getPieceByCoordinate(7, 7)

	bPawn.move(&Square{x: 7, y: 5})
	if bPawn.captured {
		t.Errorf("Pawn started out in captured state.")
	}

	// Should not be able to capture pieces it cannot access.
	moveErr := wQueen.move(&Square{x: 7, y: 5})
	if moveErr == nil {
		t.Errorf("Invalid move allowed for capture.")
	}

	wQueen.forceMove(&Square{x: 5, y: 3})
	wQueen.move(&Square{x: 7, y: 5})
	if !bPawn.captured {
		t.Errorf("Captured pawn not in captured state.")
	}

	if bPawn.getSquare() != nil {
		t.Errorf("Captured pawn has a non-nil position.")
	}
}
