package main

// You come at the king, you best not miss.
type King struct {
	data *PieceData
}

func (k King) move(newSquare Square) bool {
	return false
}

func (k King) pieceData() *PieceData {
	return k.data
}

func (k King) generateValidMoves(start Square) []*Square {
	// King moves one square in any direction, which seems unkingly.
	moves := []*Square{
		&Square{x: start.x + 1, y: start.y},
		&Square{x: start.x - 1, y: start.y},
		&Square{x: start.x, y: start.y + 1},
		&Square{x: start.x, y: start.y - 1},
		&Square{x: start.x + 1, y: start.y + 1},
		&Square{x: start.x + 1, y: start.y - 1},
		&Square{x: start.x - 1, y: start.y + 1},
		&Square{x: start.x - 1, y: start.y - 1},
	}
	validMoves := make([]*Square, 0)

	for _, move := range moves {
		if !move.isValid() {
			continue
		}
		status := k.pieceData().evaluateSquare(move)
		if status == squareVacant || status == squareOccupiedByOpponent {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}

func (k King) getShorthand() string {
	return "K"
}
