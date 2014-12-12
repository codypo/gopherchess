package main

const startSquare = 1
const endSquare = 8
const numPlayers = 2
const numPiecesPerPlayer = 16

type Color int

const (
	White Color = 1 << iota
	Black
	Undefined
)
