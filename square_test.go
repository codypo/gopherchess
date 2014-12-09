package main

import "testing"

func TestSquareIsValid(t *testing.T) {
	for i := -10; i < 10; i++ {
		s := Square{x: i, y: i}
		shouldBeValid := i >= 0 && i <= 7
		if shouldBeValid && !s.isValid() {
			t.Errorf("%d, %d is not a valid square.", i, i)
		} else if !shouldBeValid && s.isValid() {
			t.Errorf("%d, %d is a valid square", i, i)
		}
	}

	/*s := Square{x: -1, y: -1}
	if s.isValid() {
		t.Error("Expected invalid.")
	}

	s := Square(x:*/
}
