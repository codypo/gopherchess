package main

// You come at the king, you best not miss.
type King struct {
	piece *Piece
}

func (k King) getPiece() *Piece {
	return k.piece
}

func (k King) generateMoves(start Square) []*Square {
	// TODO: Castling!

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
		status := k.getPiece().evaluateSquare(move)
		if status == squareVacant || status == squareOccupiedByOpponent {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}

func (k King) getShorthand() string {
	return "K"
}
