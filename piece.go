package main

type Piece interface {
	move() bool
	pieceData() PieceData
}

type PieceData struct {
	color    Color
	square   Square
	captured bool
}
