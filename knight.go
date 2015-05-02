package main

type Knight struct {
	piece *Piece
}

func (k *Knight) setPiece(piece *Piece) {
	k.piece = piece
}

func (k Knight) getPiece() *Piece {
	return k.piece
}

func (k Knight) generateMoves(start Square) []*Square {
	return k.getPiece().generateKnightMoves(start)
}

func (k Knight) getShorthand() string {
	return "N"
}
