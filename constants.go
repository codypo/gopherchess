package main

const (
	startSquare        = 1
	endSquare          = 8
	numPlayers         = 2
	numPiecesPerPlayer = 16
	// Offset between 0 and a in ascii table, offset by 1 for
	// algebraic notation.  Magic!
	asciiOffsetForX = 48

	UnknownPieceType = '?'
)

type PieceType byte

// Slightly weird naming so to as not conflict with struct names.
const (
	PawnType   PieceType = 'p'
	RookType             = 'R'
	BishopType           = 'B'
	KnightType           = 'N'
	QueenType            = 'Q'
	KingType             = 'K'
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
