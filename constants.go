package main

const (
	startSquare        = 1
	endSquare          = 8
	numPlayers         = 2
	numPiecesPerPlayer = 16
)

type PieceType int

// Slightly weird naming so to as not conflict with struct names.
const (
	PawnType PieceType = iota
	RookType
	BishopType
	KnightType
	QueenType
	KingType
)

type Color int

const (
	White Color = iota
	Black
)

type SquareState int

const (
	SquareInvalid SquareState = iota
	SquareVacant
	SquareOccupiedByMe
	SquareOccupiedByOpponent
)

type GameState int

const (
	GameOn GameState = iota
	WhiteInCheck
	WhiteCheckmated
	BlackInCheck
	BlackCheckmated
	Draw
)
