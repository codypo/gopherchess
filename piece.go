package main

type Piece interface {
	move() bool
}

type PieceData struct {
	color    Color
	square   Square
	captured bool
}
