package main

type Piece interface {
	move(newSquare Square) bool
	pieceData() *PieceData
	generateValidMoves() []*Square
}

// Piece data exists because there's no concept of an abstract
// class in go.  So, we have the piece interface and the piece
// data struct.
type PieceData struct {
	color    Color
	captured bool
	moves    []*Square
	board    *Board
}

func (p PieceData) getSquare() *Square {
	return p.moves[len(p.moves)-1]
}

func NewPieceData(color Color, square *Square, board *Board) *PieceData {
	pd := new(PieceData)
	pd.color = color
	pd.captured = false
	pd.moves = []*Square{square}
	pd.board = board
	return pd
}
