package main

type Knight struct {
	piece *Piece
}

func (k *Knight) setPiece(piece *Piece) {
	k.piece = piece
}

func (k Knight) getPiece() *Piece {
	return k.piece
}

func (k Knight) generateMoves(start Square) []*Square {
	// Unfortunately, knight moves have nothing to do with that Bob Seger song.
	// Knights can move 2 spots vertically and 1 horizontally, or 2 horizontally
	// and 1 vertically.
	moves := []*Square{
		&Square{x: start.x + 2, y: start.y + 1},
		&Square{x: start.x + 2, y: start.y - 1},
		&Square{x: start.x - 2, y: start.y + 1},
		&Square{x: start.x - 2, y: start.y - 1},
		&Square{x: start.x + 1, y: start.y + 2},
		&Square{x: start.x + 1, y: start.y - 2},
		&Square{x: start.x - 1, y: start.y + 2},
		&Square{x: start.x - 1, y: start.y - 2},
	}

	validMoves := make([]*Square, 0)

	for _, move := range moves {
		status := k.getPiece().evaluateSquare(move)
		if status == squareVacant || status == squareOccupiedByOpponent {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}

func (k Knight) getShorthand() string {
	return "k"
}
