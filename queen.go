package main

type Queen struct {
	piece *Piece
}

func (q *Queen) setPiece(piece *Piece) {
	q.piece = piece
}

func (q Queen) getPiece() *Piece {
	return q.piece
}

func (q Queen) generateMoves(start Square) []*Square {
	// Queens are mighty can moth vertically, horizontally,
	// and diagonally.
	diagMoves := q.getPiece().generateDiagonalMoves(start)
	straightMoves := q.getPiece().generateStraightMoves(start)
	return append(diagMoves, straightMoves...)
}
