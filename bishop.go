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

func (b Bishop) generateValidMoves(start Square) []*Square {
	moves := make([]*Square, 0)

	proceedUpLeft := true
	proceedDownLeft := true
	proceedUpRight := true
	proceedDownRight := true

	for i := startSquare; i <= endSquare; i++ {
		// Evaluate the next move up and to the left.
		if proceedUpLeft {
			move := &Square{x: start.x - i, y: start.y + i}
			status := b.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			// You can only proceed onward if you're looking at a vacant square.
			proceedUpLeft = (status == squareVacant)
		}

		// Evaluate the next move down and to the left.
		if proceedDownLeft {
			move := &Square{x: start.x - i, y: start.y - i}
			status := b.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			proceedDownLeft = (status == squareVacant)
		}

		// Evaluate the next move up and to the right.
		if proceedUpRight {
			move := &Square{x: start.x + i, y: start.y + i}
			status := b.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedUpRight = (status == squareVacant)
		}

		// Evaluate the next move down and to the right
		if proceedDownRight {
			move := &Square{x: start.x + i, y: start.y - i}
			status := b.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedDownRight = (status == squareVacant)
		}
	}

	return moves
}
