package main

type Piece interface {
	move(newSquare Square) bool
	pieceData() *PieceData
	generateValidMoves(start Square) []*Square
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

func (p PieceData) addMove(square *Square) {
	p.moves = append(p.moves, square)
}

func (p PieceData) matchesCoordinates(x int, y int) bool {
	return (p.getSquare().x == x) && (p.getSquare().y == y)
}

func (p PieceData) y() int {
	return p.getSquare().y
}

func (p PieceData) x() int {
	return p.getSquare().x
}

func (p PieceData) evaluateSquare(square *Square) int {
	return p.board.evaluateSquare(p.color, square)
}

func NewPieceData(color Color, square *Square, board *Board) *PieceData {
	pd := new(PieceData)
	pd.color = color
	pd.captured = false
	pd.moves = []*Square{square}
	pd.board = board
	return pd
}
