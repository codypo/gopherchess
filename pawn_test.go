package main

import (
	"testing"
)

func TestPawnGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	pawn1 := b.getPieceByCoordinates(2, 2)

	// Pawns aren't immediately blocked by other pieces. Go nuts, pawns!
	genMoves := pawn1.generateMoves(Square{x: 2, y: 2})
	startMoves := []*Square{
		&Square{x: 2, y: 3},
		&Square{x: 2, y: 4},
	}

	arraysMatch, err := squaresMatch(genMoves, startMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}

	// Go ahead and move the pawn one spot forward.  With its first move made,
	// it can no longer move forward 2 spots.
	pawn1.move(startMoves[0])

	secondMoves := []*Square{&Square{x: 2, y: 4}}
	genMoves = pawn1.generateMoves(*pawn1.getSquare())

	arraysMatch, err = squaresMatch(genMoves, secondMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}

func TestPawnCanCapture(t *testing.T) {
	b := NewBoard()
	wPawn := b.getPieceByCoordinates(2, 2)
	bPawn := b.getPieceByCoordinates(7, 7)

	bPawn.move(&Square{x: 7, y: 5})
	if bPawn.captured {
		t.Errorf("Pawn started out in captured state.")
	}

	// Should not be able to capture pieces it cannot access.
	moveErr := wPawn.move(&Square{x: 7, y: 5})
	if moveErr == nil {
		t.Errorf("Invalid move allowed for capture.")
	}

	wPawn.forceMove(&Square{x: 6, y: 4})
	wPawn.move(&Square{x: 7, y: 5})
	if !bPawn.captured {
		t.Errorf("Captured pawn not in captured state.")
	}

	if bPawn.getSquare() != nil {
		t.Errorf("Captured pawn has a non-nil position.")
	}
}
