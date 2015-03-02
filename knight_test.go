package main

import (
	"testing"
)

func TestKnightGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	knight1, _ := white.getPieceByCoordinate(2, 1)

	// Knights are tricksy and can move immediately.
	genMoves := knight1.generateMoves(Square{x: 2, y: 1})
	startMoves := []*Square{
		&Square{x: 1, y: 3},
		&Square{x: 3, y: 3},
	}

	arraysMatch, err := squaresMatch(genMoves, startMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}

	// Sneakiness to force a move without validation.
	genMoves = knight1.generateMoves(Square{x: 4, y: 4})

	// We selected a starting square so we can move up left, up right, down left, down right.
	expMoves := []*Square{
		&Square{x: 3, y: 6},
		&Square{x: 5, y: 6},
		&Square{x: 6, y: 5},
		&Square{x: 6, y: 3},
		&Square{x: 2, y: 5},
		&Square{x: 2, y: 3},
	}

	arraysMatch, err = squaresMatch(genMoves, expMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}

func TestKnightCanCapture(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)
	wKnight, _ := white.getPieceByCoordinate(2, 1)
	bPawn, _ := black.getPieceByCoordinate(7, 7)

	bPawn.move(&Square{x: 7, y: 5})
	if bPawn.captured {
		t.Errorf("Pawn started out in captured state.")
	}

	// Should not be able to capture pieces it cannot access.
	moveErr := wKnight.move(&Square{x: 7, y: 5})
	if moveErr == nil {
		t.Errorf("Invalid move allowed for capture.")
	}

	wKnight.forceMove(&Square{x: 6, y: 3})
	wKnight.move(&Square{x: 7, y: 5})
	if !bPawn.captured {
		t.Errorf("Captured pawn not in captured state.")
	}

	if bPawn.getSquare() != nil {
		t.Errorf("Captured pawn has a non-nil position.")
	}
}
