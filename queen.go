package main

type Queen struct {
	data *PieceData
}

func (q Queen) move(newSquare Square) bool {
	return false
}

func (q Queen) pieceData() *PieceData {
	return q.data
}

func (q Queen) generateMoves(start Square) []*Square {
	diagMoves := q.pieceData().generateDiagonalMoves(start)
	straightMoves := q.pieceData().generateStraightMoves(start)
	return append(diagMoves, straightMoves...)
}

func (q Queen) getShorthand() string {
	return "q"
}
