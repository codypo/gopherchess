package main

type Queen struct {
	data *PieceData
}

func (q Queen) move() bool {
	return false
}

func (q Queen) pieceData() *PieceData {
	return q.data
}
