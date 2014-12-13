package main

type Bishop struct {
	data PieceData
}

func (b Bishop) move() bool {
	return false
}

func (b Bishop) pieceData() PieceData {
	return b.data
}
