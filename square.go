package main

import (
	"strconv"
)

// Represents a square of the board.
type Square struct {
	x int
	y int
}

// Is the struct in question valid?
func (s Square) isValid() bool {
	return (s.x >= startSquare && s.x <= endSquare) && (s.y >= startSquare && s.y <= endSquare)
}

// Get the color of a square.
func (s Square) color() (color Color) {
	if s.x%2 == 0 && s.y%2 == 0 {
		return White
	}
	return Black
}

// Compares two squares.
func (s1 Square) equals(s2 Square) bool {
	return (s1.x == s2.x) && (s1.y == s2.y)
}

// Determines if the notation for a square is valid.
func isSquareNotationValid(square string) bool {
	// Convert 'a1' into x: 1, y: 1

	// First, turn something 'a' into an int.
	x := int(square[0] - '0')
	x = x - asciiOffsetForX

	// Atoi is shorthand for ParseInt.  Naming is not so intuitive.
	y, err := strconv.Atoi(square[1:2])
	if err != nil {
		return false
	}
	return Square{x: x, y: y}.isValid()
}
