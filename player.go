package main

import "fmt"

type Player struct {
	name   string
	color  Color
	pieces [numPiecesPerPlayer]*Piece
}

func NewPlayer(color Color, board *Board) *Player {
	p := new(Player)
	p.color = color

	// Piece initialization goes here.
	pieceIndex := 0

	// p p p p p p p p.
	// r k b q k b k r.
	pawnRow := startSquare + 1
	rookRow := startSquare
	if color == Black {
		pawnRow = endSquare - 1
		rookRow = endSquare
	}

	// Initialize a square and piece data instance we can reuse.
	s := &Square{x: startSquare, y: startSquare}
	pawn := NewPiece(color, s, board, PawnType)

	// Populate pawns.
	for x := startSquare; x <= endSquare; x++ {
		s = &Square{x: x, y: pawnRow}
		pawn = NewPiece(color, s, board, PawnType)
		p.pieces[pieceIndex] = pawn
		pieceIndex++
	}

	// Populate rooks.
	s = &Square{x: startSquare, y: rookRow}
	rook := NewPiece(color, s, board, RookType)
	p.pieces[pieceIndex] = rook
	pieceIndex++

	// TODO: This pieceIndex part is silly.
	s = &Square{x: endSquare, y: rookRow}
	rook = NewPiece(color, s, board, RookType)
	p.pieces[pieceIndex] = rook
	pieceIndex++

	// Populate knights.
	s = &Square{x: startSquare + 1, y: rookRow}
	knight := NewPiece(color, s, board, KnightType)
	p.pieces[pieceIndex] = knight
	pieceIndex++

	s = &Square{x: endSquare - 1, y: rookRow}
	knight = NewPiece(color, s, board, KnightType)
	p.pieces[pieceIndex] = knight
	pieceIndex++

	// Populate bishops.
	s = &Square{x: startSquare + 2, y: rookRow}
	bishop := NewPiece(color, s, board, BishopType)
	p.pieces[pieceIndex] = bishop
	pieceIndex++

	s = &Square{x: endSquare - 2, y: rookRow}
	bishop = NewPiece(color, s, board, BishopType)
	p.pieces[pieceIndex] = bishop
	pieceIndex++

	// Populate the queen.
	s = &Square{x: startSquare + 3, y: rookRow}
	queen := NewPiece(color, s, board, QueenType)
	p.pieces[pieceIndex] = queen
	pieceIndex++

	// Populate the king.
	s = &Square{x: startSquare + 4, y: rookRow}
	king := NewPiece(color, s, board, KingType)
	p.pieces[pieceIndex] = king
	pieceIndex++

	return p
}

func (player Player) getPieceByCoordinate(x int, y int) (*Piece, error) {
	for _, piece := range player.pieces {
		if piece.matchesCoordinates(x, y) {
			return piece, nil
		}
	}

	return nil, fmt.Errorf("No piece found at %d, %d", x, y)
}

func (p Player) printPieces() {
	for _, piece := range p.pieces {
		fmt.Printf("Piece coords %d, %d\n", piece.getSquare().x, piece.getSquare().y)
	}
}

// Gets the king piece for this player.  Used to determine if player
// is in check.
func (p Player) getKing() (*Piece, error) {
	for _, p := range p.pieces {
		// TODO: I really hate the PieceType naming.
		if p.pieceType == KingType {
			return p, nil
		}
	}

	return nil, fmt.Errorf("No king found.  This very, very odd.")
}

// Can any piece for this player move to a given square?  Used to
// determine if player is in check.
func (p Player) canMoveToSquare(s Square) bool {
	for _, piece := range p.pieces {
		if piece.captured {
			continue
		}

		if piece.canMoveToSquare(s) {
			return true
		}
	}

	return false
}

// Performs a deep copy of a player's pieces.
func (p Player) deepCopy(board *Board, color Color) *Player {
	c := new(Player)
	c.color = color
	for i, piece := range p.pieces {
		c.pieces[i] = piece.deepCopy(board)
	}

	// TODO: Let's drop this and move it all into board.
	return c
}
