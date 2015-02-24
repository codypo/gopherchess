package main

import (
	"testing"
)

func TestKnightGeneratesValidMoves(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	knight1 := white.getPieceByCoordinate(2, 1)

	// Knights are tricksy and can move immediately.
	genMoves := knight1.generateValidMoves(Square{x: 2, y: 1})
	var startMoves [2]*Square
	startMoves[0] = &Square{x: 1, y: 3}
	startMoves[1] = &Square{x: 3, y: 3}

	arraysMatch, err := squareArraysMatch(genMoves, startMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}

	// Sneakiness to force a move without validation.
	genMoves = knight1.generateValidMoves(Square{x: 4, y: 4})

	var expMoves [6]*Square
	// We selected a starting square so we can move up left, up right, down left, down right.
	expMoves[0] = &Square{x: 3, y: 6}
	expMoves[1] = &Square{x: 5, y: 6}
	expMoves[2] = &Square{x: 6, y: 5}
	expMoves[3] = &Square{x: 6, y: 3}
	expMoves[4] = &Square{x: 3, y: 2}
	expMoves[5] = &Square{x: 5, y: 2}

	arraysMatch, err = squareArraysMatch(genMoves, expMoves[0:])
	if !arraysMatch {
		t.Errorf(err.Error())
	}
}
