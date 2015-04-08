package main

import (
	"fmt"
	"testing"
)

func TestRookGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	//b.prettyPrint()
	//b.dumpSquares()
	rook1 := b.getPieceByCoordinates(1, 1)

	// Initially, the white rook can't move because it's boxed in.
	moves := rook1.generateMoves(Square{x: 1, y: 1})
	if len(moves) > 0 {
		t.Errorf("Rook in starting position has valid moves.")
	}

	// Sneakiness to force a move without validation.
	genMoves := rook1.generateMoves(Square{x: 4, y: 4})

	// Rook can move vertically up to the first opponent piece.
	// Rook can move horizontally across the board.
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
	}

	arraysMatch, err := squaresMatch(genMoves, expMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}

func TestRookCanCapture(t *testing.T) {
	b := NewBoard()
	wRook := b.getPieceByCoordinates(8, 1)
	bPawn := b.getPieceByCoordinates(7, 7)
	fmt.Printf("1 Can Capture\n")

	bPawn.move(&Square{x: 7, y: 5})
	if bPawn.captured {
		t.Errorf("Pawn started out in captured state.")
	}
	fmt.Printf("1.1 Can Capture\n")
	b.prettyPrint()

	// Should not be able to capture pieces it cannot access.
	moveErr := wRook.move(&Square{x: 7, y: 5})
	if moveErr == nil {
		t.Errorf("Invalid move allowed for capture.")
	}
	fmt.Printf("1.2 Can Capture\n")

	wRook.forceMove(&Square{x: 7, y: 3})
	b.prettyPrint()
	fmt.Printf("1.3 Can Capture\n")
	wRook.move(&Square{x: 7, y: 5})
	b.prettyPrint()
	if !bPawn.captured {
		t.Errorf("Captured pawn not in captured state.")
	}
	fmt.Printf("1.3 Can Capture\n")

	if bPawn.getSquare() != nil {
		t.Errorf("Captured pawn has a non-nil position.")
	}
}

func TestRookCannotMoveAndLeadToOwnCheck(t *testing.T) {
	b := NewBoard()

	wRook := b.getPieceByCoordinates(8, 1)
	wKing := b.getKing(White)
	bRook := b.getPieceByCoordinates(8, 8)

	// Black rook would have white King in check if white
	// rook just moved out of the way.  White rook can't
	// perform that moev, though.

	bRook.forceMove(&Square{x: 8, y: 6})
	wRook.forceMove(&Square{x: 8, y: 5})
	wKing.forceMove(&Square{x: 8, y: 4})

	// For the love of Jeebus, stay there, white rook!
	moveErr := wRook.move(&Square{x: 1, y: 5})
	if moveErr == nil {
		t.Errorf("White rook should not be able to move and leave white king in check.")
	}

}
