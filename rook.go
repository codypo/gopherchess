package main

type Rook struct {
	data PieceData
}

func (r Rook) move() bool {
	return false
}

func (r Rook) pieceData() PieceData {
	return r.data
}
