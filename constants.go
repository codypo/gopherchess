package main

const startSquareIndex = 1
const numSquaresWide = 8
const numSquaresTall = numSquaresWide
const numPlayers = 2
const numPiecesPerPlayer = 16

type Color int

const (
	White Color = 1 << iota
	Black
	Undefined
)
