package main

const (
	startSquare        = 1
	endSquare          = 8
	numPlayers         = 2
	numPiecesPerPlayer = 16
)

// Slightly weird naming so to as not conflict with struct names.
type PieceType int

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
	Default GameState = 1 << iota // TODO: Don't like the name Normal.
	WhiteInCheck
	WhiteInCheckMate
	BlackInCheck
	BlackInCheckMate
)
