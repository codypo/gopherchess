package main

type Queen struct {
	data *PieceData
}

func (q Queen) move(newSquare Square) bool {
	return false
}

func (q Queen) pieceData() *PieceData {
	return q.data
}

func (q Queen) generateValidMoves(start Square) []*Square {
	return make([]*Square, 10, 10)
}
