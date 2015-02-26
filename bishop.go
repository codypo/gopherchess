package main

type Bishop struct {
	data *PieceData
}

func (b Bishop) move(newSquare Square) bool {
	return false
}

func (b Bishop) pieceData() *PieceData {
	return b.data
}

func (b Bishop) generateMoves(start Square) []*Square {
	return b.pieceData().generateDiagonalMoves(start)
}

func (b Bishop) getShorthand() string {
	return "b"
}
