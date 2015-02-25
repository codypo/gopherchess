package main

const (
	startSquare        = 1
	endSquare          = 8
	numPlayers         = 2
	numPiecesPerPlayer = 16

	squareInvalid            = -1
	squareVacant             = 0
	squareOccupiedByMe       = 1
	squareOccupiedByOpponent = 128 // Random witchcraft.
)

type Color int

const (
	White Color = 1 << iota
	Black
	Undefined
)
