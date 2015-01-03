package main

const startSquare = 1
const endSquare = 8
const numPlayers = 2
const numPiecesPerPlayer = 16

const squareInvalid = -1
const squareVacant = 0
const squareOccupiedByMe = 1
const squareOccupiedByOpponent = 128 // Random witchcraft.

type Color int

const (
	White Color = 1 << iota
	Black
	Undefined
)
