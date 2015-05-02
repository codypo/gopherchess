package main

import "testing"

func TestPawnsPopulatedCorrectly(t *testing.T) {
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
		c := s.color()

		if s.isValid() {
			// Even squares are light.
			if i%2 == 0 && c != White {
				t.Errorf("%d, %d is not white.", i, i)
			}

			// Odd squares are dark.
			if i%2 != 0 && c != Black {
				t.Errorf("%d, %d is not black.", i, i)
			}
		}
	}
}
