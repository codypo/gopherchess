package main

// You come at the king, you best not miss.
type King struct {
	data *PieceData
}

func (k King) move(newSquare Square) bool {
	return false
}

func (k King) pieceData() *PieceData {
	return k.data
}

func (k King) generateValidMoves(start Square) []*Square {
	return make([]*Square, 10, 10)
}
