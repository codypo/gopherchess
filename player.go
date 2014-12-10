package main

type Player struct {
	name   string
	color  Color
	pieces []Piece
}

func NewPlayer(color Color) *Player {
	p := new(Player)
	p.color = color

	// Piece initialization goes here.

	return p
}
