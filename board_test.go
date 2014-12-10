package main

import "testing"

func TestBoardNew(t *testing.T) {
	b := NewBoard()
	totalSquares := 0
	for i, _ := range b.squares {
		for j, _ := range b.squares[i] {
			totalSquares++

			s := b.squares[i][j]
			if !s.isValid() {
				t.Errorf("Board contains invalid square at %d, %d", s.x, s.y)
			}
		}
	}
	if totalSquares != 64 {
		t.Errorf("Board contains %d squares.", totalSquares)
	}

	if len(b.players) != 2 {
		t.Errorf("Board contains %d players.", len(b.players))
	}
	if b.players[0].color != White {
		t.Errorf("Player 0 is not White.")
	}
	if b.players[1].color != Black {
		t.Errorf("Player 1 is not Black.")
	}
}
