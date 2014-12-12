package main

import (
	"errors"
)

type Square struct {
	x int
	y int
}

func (s Square) isValid() bool {
	return (s.x >= startSquare && s.x <= endSquare) && (s.y >= startSquare && s.y <= endSquare)
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
