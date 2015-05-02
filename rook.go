package main

type Rook struct {
	piece *Piece
}

func (r *Rook) setPiece(piece *Piece) {
	r.piece = piece
}

func (r Rook) getPiece() *Piece {
	return r.piece
}

func (r Rook) generateMoves(start Square) []*Square {
	return r.getPiece().generateStraightMoves(start)
}

func (r Rook) getShorthand() string {
	return "R"
}
