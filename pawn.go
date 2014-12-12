package main

type Pawn struct {
	data PieceData
}

func (p Pawn) move() bool {
	return false
}

func (p Pawn) pieceData() PieceData {
	return p.data
}
