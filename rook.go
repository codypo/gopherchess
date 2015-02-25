package main

type Rook struct {
	data *PieceData
}

func (r Rook) move(newSquare Square) bool {
	return false
}

func (r Rook) pieceData() *PieceData {
	return r.data
}

func (r Rook) generateValidMoves(start Square) []*Square {
	return r.pieceData().generateStraightMoves(start)
}

func (r Rook) getShorthand() string {
	return "r"
}
