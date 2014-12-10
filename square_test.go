package main

import "testing"

func TestSquareIsValid(t *testing.T) {
	for i := -10; i < 10; i++ {
		s := Square{x: i, y: i}
		shouldBeValid := i >= 1 && i <= 8
		if shouldBeValid && !s.isValid() {
			t.Errorf("%d, %d is not a valid square.", i, i)
		} else if !shouldBeValid && s.isValid() {
			t.Errorf("%d, %d is a valid square", i, i)
		}
	}
}

func TestSquareColor(t *testing.T) {
	for i := -10; i < 10; i++ {
		s := Square{x: i, y: i}
		c, err := s.color()

		if s.isValid() {
			// Color should only return an error if the square is invalid.
			if err != nil || c == Undefined {
				t.Errorf("%d, %d is a valid square, but returning a bad color.", i, i)
			}

			// Even squares are light.
			if i%2 == 0 && c != White {
				t.Errorf("%d, %d is not white.", i, i)
			}

			// Odd squares are dark.
			if i%2 != 0 && c != Black {
				t.Errorf("%d, %d is not black.", i, i)
			}
		} else {
			if c != Undefined {
				t.Errorf("Invalid square %d, %d has a valid color.", i, i)
			}
		}
	}
}
