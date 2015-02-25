package main

type Pawn struct {
	data *PieceData
}

func (p Pawn) move(newSquare Square) bool {
	return false
}

func (p Pawn) pieceData() *PieceData {
	return p.data
}

func (p Pawn) generateValidMoves(start Square) []*Square {
	return make([]*Square, 10, 10)
}

func (p Pawn) getShorthand() string {
	return "p"
}
