package main

import (
	"testing"
)

func TestPawnGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	pawn1, _ := white.getPieceByCoordinate(2, 2)

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
