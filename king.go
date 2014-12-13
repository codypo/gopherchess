package main

type King struct {
	data PieceData
}

func (k King) move() bool {
	return false
}

func (k King) pieceData() PieceData {
	return k.data
}
