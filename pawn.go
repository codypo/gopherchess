package main

type Pawn struct {
	piece *Piece
}

func (p *Pawn) setPiece(piece *Piece) {
	p.piece = piece
}

func (p Pawn) getPiece() *Piece {
	return p.piece
}

func (p Pawn) generateMoves(start Square) []*Square {
	// Pawn can move vertically 1 square.
	// If it's the pawn's first move, he can jump 2 squares.

	// TODO: en passant and promotion.
	moveDirection := 1
	if p.piece.color == Black {
		moveDirection = -1
	}

	moves := make([]*Square, 1)
	moves[0] = &Square{x: start.x, y: start.y + (1 * moveDirection)}

	if len(p.piece.moves) == 1 {
		moves = append(moves, &Square{x: start.x, y: start.y + (2 * moveDirection)})
	}

	// TODO: when we add capturing, just chop this and validate the moves
	// one by one, since the capture rules are weird.

	validMoves := make([]*Square, 0)
	for _, move := range moves {
		status := p.getPiece().evaluateSquare(move)
		if status == SquareVacant {
			validMoves = append(validMoves, move)
		}
	}

	// Certain moves are only valid if we can capture an opponent.
	// In this case, we make a special case for a pawn's diagonal move.
	captureLeft := &Square{x: start.x - 1, y: start.y + (1 * moveDirection)}
	if p.getPiece().evaluateSquare(captureLeft) == SquareOccupiedByOpponent {
		validMoves = append(validMoves, captureLeft)
	}
	captureRight := &Square{x: start.x + 1, y: start.y + (1 * moveDirection)}
	if p.getPiece().evaluateSquare(captureRight) == SquareOccupiedByOpponent {
		validMoves = append(validMoves, captureRight)
	}

	return validMoves
}

func (p Pawn) getShorthand() string {
	return "p"
}
