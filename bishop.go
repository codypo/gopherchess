package main

type Bishop struct {
	piece *Piece
}

func (b Bishop) getPiece() *Piece {
	return b.piece
}

func (b Bishop) generateMoves(start Square) []*Square {
	// Bishops can only move diagonally.
	return b.getPiece().generateDiagonalMoves(start)
}

func (b Bishop) getShorthand() string {
	return "b"
}
