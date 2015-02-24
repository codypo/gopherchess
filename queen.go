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

func (q Queen) generateValidMoves(start Square) []*Square {
	moves := make([]*Square, 0)

	proceedUp := true
	proceedDown := true
	proceedLeft := true
	proceedRight := true

	proceedUpLeft := true
	proceedDownLeft := true
	proceedUpRight := true
	proceedDownRight := true

	for i := startSquare; i <= endSquare; i++ {
		// Evaluate the next move up.
		if proceedUp {
			move := &Square{x: start.x, y: start.y + i}
			status := q.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			// You can only proceed onward if you're looking at a vacant square.
			proceedUp = (status == squareVacant)
		}

		// Evaluate the next move down.
		if proceedDown {
			move := &Square{x: start.x, y: start.y - i}
			status := q.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			proceedDown = (status == squareVacant)
		}

		// Evaluate the next move right.
		if proceedRight {
			move := &Square{x: start.x + i, y: start.y}
			status := q.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedRight = (status == squareVacant)
		}

		// Evaluate the next move left.
		if proceedLeft {
			move := &Square{x: start.x - i, y: start.y}
			status := q.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedLeft = (status == squareVacant)
		}

		// Evaluate the next move up and to the left.
		if proceedUpLeft {
			move := &Square{x: start.x - i, y: start.y + i}
			status := q.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			// You can only proceed onward if you're looking at a vacant square.
			proceedUpLeft = (status == squareVacant)
		}

		// Evaluate the next move down and to the left.
		if proceedDownLeft {
			move := &Square{x: start.x - i, y: start.y - i}
			status := q.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			proceedDownLeft = (status == squareVacant)
		}

		// Evaluate the next move up and to the right.
		if proceedUpRight {
			move := &Square{x: start.x + i, y: start.y + i}
			status := q.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedUpRight = (status == squareVacant)
		}

		// Evaluate the next move down and to the right
		if proceedDownRight {
			move := &Square{x: start.x + i, y: start.y - i}
			status := q.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedDownRight = (status == squareVacant)
		}
	}

	return moves
}
