package main

import "testing"

// TODO: All thsi needs to go through new board.

func TestPlayerNewPopulatesPawns(t *testing.T) {
	b := NewBoard()

	p := b.getPlayer(White)
	if len(p.pieces) != 16 {
		t.Errorf("Player has %d pieces.", len(p.pieces))
	}

	for i := 0; i < 8; i++ {
		pawn := p.pieces[i]
		p_loc := pawn.getSquare()
		if p_loc.x != i+1 || p_loc.y != 2 {
			t.Errorf("Pawn at index %d is in.getSquare() %d, %d", i, p_loc.x, p_loc.y)
		}
	}

	// Verify that the black player's pawns exist.
	p = b.getPlayer(Black)

	for i := 0; i < 8; i++ {
		pawn := p.pieces[i]
		p_loc := pawn.getSquare()
		if p_loc.x != i+1 || p_loc.y != 7 {
			t.Errorf("Pawn at index %d is in.getSquare() %d, %d", i, p_loc.x, p_loc.y)
		}
	}
}

func TestPlayerNewPopulatesRooks(t *testing.T) {
	b := NewBoard()
	p := b.getPlayer(White)

	// First 8 pieces are pawns.  2 immediately subsequent rooks.
	rookIndex := 8
	rook := p.pieces[rookIndex]
	r_loc := rook.getSquare()
	if r_loc.x != 1 || r_loc.y != 1 {
		t.Errorf("Rook 1 is located at %d, %d", r_loc.x, r_loc.y)
	}

	rook = p.pieces[rookIndex+1]
	r_loc = rook.getSquare()
	if r_loc.x != 8 || r_loc.y != 1 {
		t.Errorf("Rook 2 located at %d, %d", r_loc.x, r_loc.y)
	}

	// Now verify that the black player's rooks exist.
	p = b.getPlayer(Black)

	rook = p.pieces[rookIndex]
	r_loc = rook.getSquare()
	if r_loc.x != 1 || r_loc.y != 8 {
		t.Errorf("Rook 3 is located at %d, %d", r_loc.x, r_loc.y)
	}

	rook = p.pieces[rookIndex+1]
	r_loc = rook.getSquare()
	if r_loc.x != 8 || r_loc.y != 8 {
		t.Errorf("Rook 4 located at %d, %d", r_loc.x, r_loc.y)
	}
}

func TestPlayerNewPopulatesKnights(t *testing.T) {
	b := NewBoard()
	p := b.getPlayer(White)

	// First 8 pieces are pawns.  Then rooks, then knights.
	knightIndex := 10
	knight := p.pieces[knightIndex]
	k_loc := knight.getSquare()
	if k_loc.x != 2 || k_loc.y != 1 {
		t.Errorf("Knight 1 is located at %d, %d", k_loc.x, k_loc.y)
	}

	knight = p.pieces[knightIndex+1]
	k_loc = knight.getSquare()
	if k_loc.x != 7 || k_loc.y != 1 {
		t.Errorf("Knight 2 located at %d, %d", k_loc.x, k_loc.y)
	}

	// Now verify the existence of the black knights.
	p = b.getPlayer(Black)

	knight = p.pieces[knightIndex]
	k_loc = knight.getSquare()
	if k_loc.x != 2 || k_loc.y != 8 {
		t.Errorf("Knight 1 is located at %d, %d", k_loc.x, k_loc.y)
	}

	knight = p.pieces[knightIndex+1]
	k_loc = knight.getSquare()
	if k_loc.x != 7 || k_loc.y != 8 {
		t.Errorf("Knight 2 located at %d, %d", k_loc.x, k_loc.y)
	}
}

func TestPlayerNewPopulatesBishops(t *testing.T) {
	b := NewBoard()
	p := b.getPlayer(White)

	// First 8 pieces are pawns.  Then rooks, then knights, then bishops.
	bishopIndex := 12
	bishop := p.pieces[bishopIndex]
	b_loc := bishop.getSquare()
	if b_loc.x != 3 || b_loc.y != 1 {
		t.Errorf("Bishop 1 is located at %d, %d", b_loc.x, b_loc.y)
	}

	bishop = p.pieces[bishopIndex+1]
	b_loc = bishop.getSquare()
	if b_loc.x != 6 || b_loc.y != 1 {
		t.Errorf("Bishop 2 located at %d, %d", b_loc.x, b_loc.y)
	}

	// Now verify the existence of the black bishops.
	p = b.getPlayer(Black)

	bishop = p.pieces[bishopIndex]
	b_loc = bishop.getSquare()
	if b_loc.x != 3 || b_loc.y != 8 {
		t.Errorf("Bishop 1 is located at %d, %d", b_loc.x, b_loc.y)
	}

	bishop = p.pieces[bishopIndex+1]
	b_loc = bishop.getSquare()
	if b_loc.x != 6 || b_loc.y != 8 {
		t.Errorf("Bishop 2 located at %d, %d", b_loc.x, b_loc.y)
	}
}

func TestPlayerNewPopulatesQueen(t *testing.T) {
	b := NewBoard()
	p := b.getPlayer(White)

	// First 8 pieces are pawns.  Then rooks, then knights, then bishops, then queen.
	queenIndex := 14
	queen := p.pieces[queenIndex]
	q_loc := queen.getSquare()
	if q_loc.x != 4 || q_loc.y != 1 {
		t.Errorf("Queen is located at %d, %d", q_loc.x, q_loc.y)
	}

	// Now verify the existence of the black queen.
	p = b.getPlayer(Black)

	queen = p.pieces[queenIndex]
	q_loc = queen.getSquare()
	if q_loc.x != 4 || q_loc.y != 8 {
		t.Errorf("Queen is located at %d, %d", q_loc.x, q_loc.y)
	}
}

func TestPlayerNewPopulatesKing(t *testing.T) {
	b := NewBoard()
	p := b.getPlayer(White)

	// First 8 pieces are pawns.  Then rooks, then knights, then bishops, then queen, then king.
	kingIndex := 15
	king := p.pieces[kingIndex]
	k_loc := king.getSquare()
	if k_loc.x != 5 || k_loc.y != 1 {
		t.Errorf("King is located at %d, %d", k_loc.x, k_loc.y)
	}

	// Now verify the existence of the black king.
	p = b.getPlayer(Black)

	king = p.pieces[kingIndex]
	k_loc = king.getSquare()
	if k_loc.x != 5 || k_loc.y != 8 {
		t.Errorf("King is located at %d, %d", k_loc.x, k_loc.y)
	}
}

func BenchmarkNewPlayer(b *testing.B) {
	board := NewBoard()
	// Benchmark the performance of creating a new player, which can be
	// an allocation-heavy event.
	for n := 0; n < b.N; n++ {
		NewPlayer(White, board)
		NewPlayer(Black, board)
	}
}
