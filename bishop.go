package main

type Bishop struct {
	data *PieceData
}

func (b Bishop) move(newSquare Square) bool {
	return false
}

func (b Bishop) pieceData() *PieceData {
	return b.data
}

func (b Bishop) generateValidMoves() []*Square {
	return make([]*Square, 10, 10)
}
