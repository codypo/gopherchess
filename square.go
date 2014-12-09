package main

type Square struct {
	backgroundColor Color
	x               int
	y               int
}

func (s Square) isValid() bool {
	return (s.x >= 0 && s.x < numSquaresWide) && (s.y >= 0 && s.y < numSquaresTall)
}
