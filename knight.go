package main

type Knight struct {
	data *PieceData
}

func (k Knight) move() bool {
	return false
}

func (k Knight) pieceData() *PieceData {
	return k.data
}
