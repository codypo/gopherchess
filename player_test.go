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

	// Verify that the black player's pawns exist.
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

	// First 8 pieces are pawns.  2 immediately subsequent rooks.
	rook := p.pieces[8]
	r_loc := rook.pieceData().square
	if r_loc.x != 1 || r_loc.y != 1 {
		t.Errorf("Rook 1 is located at %d, %d", r_loc.x, r_loc.y)
	}

	rook = p.pieces[9]
	r_loc = rook.pieceData().square
	if r_loc.x != 8 || r_loc.y != 1 {
		t.Errorf("Rook 2 located at %d, %d", r_loc.x, r_loc.y)
	}

	// Now verify that the black player's rooks exist.
	p = NewPlayer(Black)

	rook = p.pieces[8]
	r_loc = rook.pieceData().square
	if r_loc.x != 1 || r_loc.y != 8 {
		t.Errorf("Rook 3 is located at %d, %d", r_loc.x, r_loc.y)
	}

	rook = p.pieces[9]
	r_loc = rook.pieceData().square
	if r_loc.x != 8 || r_loc.y != 8 {
		t.Errorf("Rook 4 located at %d, %d", r_loc.x, r_loc.y)
	}
}

func TestPlayerNewPopulatesKnights(t *testing.T) {
	p := NewPlayer(White)

	// First 8 pieces are pawns.  Then rooks, then knights.
	knight := p.pieces[10]
	k_loc := knight.pieceData().square
	if k_loc.x != 2 || k_loc.y != 1 {
		t.Errorf("Knight 1 is located at %d, %d", k_loc.x, k_loc.y)
	}

	knight = p.pieces[11]
	k_loc = knight.pieceData().square
	if k_loc.x != 7 || k_loc.y != 1 {
		t.Errorf("Knight 2 located at %d, %d", k_loc.x, k_loc.y)
	}

	// Now verify the existence of the black knights.
	p = NewPlayer(Black)

	knight = p.pieces[10]
	k_loc = knight.pieceData().square
	if k_loc.x != 2 || k_loc.y != 8 {
		t.Errorf("Knight 1 is located at %d, %d", k_loc.x, k_loc.y)
	}

	knight = p.pieces[11]
	k_loc = knight.pieceData().square
	if k_loc.x != 7 || k_loc.y != 8 {
		t.Errorf("Knight 2 located at %d, %d", k_loc.x, k_loc.y)
	}
}
