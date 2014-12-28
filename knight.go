package main

type Knight struct {
	data *PieceData
}

func (k Knight) move(newSquare Square) bool {
	return false
}

func (k Knight) pieceData() *PieceData {
	return k.data
}

func (k Knight) generateValidMoves() []*Square {
	return make([]*Square, 10, 10)
}
