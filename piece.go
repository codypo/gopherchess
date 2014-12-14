package main

type Piece interface {
	move() bool
	pieceData() *PieceData
}

// Piece data exists because there's no concept of an abstract
// class in go.  So, we have the piece interface and the piece
// data struct.
type PieceData struct {
	color    Color
	square   Square
	captured bool
	moves    []Square
}

func NewPieceData(color Color, square Square) *PieceData {
	pd := new(PieceData)
	pd.color = color
	pd.square = square
	pd.captured = false
	pd.moves = make([]Square, 10, 10)
	return pd
}
