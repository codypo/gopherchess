package main

import (
	"errors"
)

type Square struct {
	x int
	y int
}

func (s Square) isValid() bool {
	return (s.x >= 1 && s.x <= numSquaresWide) && (s.y >= 1 && s.y <= numSquaresTall)
}

func (s Square) color() (color Color, err error) {
	if !s.isValid() {
		return Undefined, errors.New("Color requested for invalid square.")
	}

	if s.x%2 == 0 && s.y%2 == 0 {
		return White, nil
	}
	return Black, nil
}
