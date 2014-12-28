package main

type Rook struct {
	data *PieceData
}

func (r Rook) move(newSquare Square) bool {
	return false
}

func (r Rook) pieceData() *PieceData {
	return r.data
}

func (r Rook) generateValidMoves() []*Square {
	return make([]*Square, 10, 10)
}
