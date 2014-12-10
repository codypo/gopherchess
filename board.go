package main

type Board struct {
	players [numPlayers]Player
	squares [numSquaresWide][numSquaresTall]Square
}

func NewBoard() *Board {
	b := new(Board)

	// TODO: is all this startSquareIndex business actually useful?

	// This feels gross, but it's a way to ensure that 0-indexed array
	// lines up with silly 1-indexed squares.
	offset := startSquareIndex - 0
	for x := startSquareIndex; x < startSquareIndex+numSquaresWide; x++ {
		for y := startSquareIndex; y < startSquareIndex+numSquaresWide; y++ {
			s := Square{x: x, y: y}
			b.squares[x-offset][y-offset] = s
		}
	}

	p0 := NewPlayer(White)
	b.players[0] = *p0

	p1 := NewPlayer(Black)
	b.players[1] = *p1

	return b
}
