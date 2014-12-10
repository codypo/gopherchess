package main

type Square struct {
	backgroundColor Color
	x               int
	y               int
}

func (s Square) isValid() bool {
	return (s.x >= 1 && s.x <= numSquaresWide) && (s.y >= 1 && s.y <= numSquaresTall)
}
