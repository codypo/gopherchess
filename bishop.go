package main

type Bishop struct {
	piece *Piece
}

func (b *Bishop) setPiece(myPiece *Piece) {
	b.piece = myPiece
}

func (b Bishop) getPiece() *Piece {
	return b.piece
}

func (b Bishop) generateMoves(start Square) []*Square {
	// Bishops can only move diagonally.
	return b.getPiece().generateDiagonalMoves(start)
}
