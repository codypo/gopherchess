package main

import "testing"

func TestPlayerNewPopulatesPawns(t *testing.T) {
	p := NewPlayer(White)
	if len(p.pieces) != 16 {
		t.Errorf("Player has %d pieces.", len(p.pieces))
	}

	for i := 0; i < 8; i++ {
		pawn := p.pieces[i]
		p_loc := pawn.pieceData().square
		if p_loc.x != i+1 || p_loc.y != 2 {
			t.Errorf("Pawn at index %d is in square %d, %d", i, p_loc.x, p_loc.y)
		}
	}

	p = NewPlayer(Black)

	for i := 0; i < 8; i++ {
		pawn := p.pieces[i]
		p_loc := pawn.pieceData().square
		if p_loc.x != i+1 || p_loc.y != 7 {
			t.Errorf("Pawn at index %d is in square %d, %d", i, p_loc.x, p_loc.y)
		}
	}
}

func TestPlayerNewPopulatesRooks(t *testing.T) {
	p := NewPlayer(White)

	// First 8 pieces are pawns.  Only 2 pawns.
	rook1 := p.pieces[8]
	r1_loc := rook1.pieceData().square
	if r1_loc.x != 1 || r1_loc.y != 1 {
		t.Errorf("Rook 1 is located at %d, %d", r1_loc.x, r1_loc.y)
	}

	rook2 := p.pieces[9]
	r2_loc := rook2.pieceData().square
	if r2_loc.x != 8 || r2_loc.y != 1 {
		t.Errorf("Rook 2 located at %d, %d", r2_loc.x, r2_loc.y)
	}
}
