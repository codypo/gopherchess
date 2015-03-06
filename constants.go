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
	PawnType PieceType = 1 << iota
	RookType
	BishopType
	KnightType
	QueenType
	KingType
)

type Color int

const (
	White Color = 1 << iota
	Black
	Undefined
)

type SquareState int

const (
	SquareInvalid SquareState = 1 << iota
	SquareVacant
	SquareOccupiedByMe
	SquareOccupiedByOpponent
)

type GameState int

const (
	GameOn GameState = 1 << iota
	WhiteInCheck
	WhiteInCheckMate
	BlackInCheck
	BlackInCheckMate
)
