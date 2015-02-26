package main

import "fmt"

type Pawn struct {
	data *PieceData
}

func (p Pawn) move(newSquare Square) bool {
	return false
}

func (p Pawn) pieceData() *PieceData {
	return p.data
}

func (p Pawn) generateMoves(start Square) []*Square {
	// Pawn can move vertically 1 square.
	// If it's the pawn's first move, he can jump 2 squares.

	// TODO: en passant
	moves := make([]*Square, 1)
	moves[0] = &Square{x: start.x, y: start.y + 1}
	fmt.Printf("GM Added first move.\n")
	fmt.Printf("GM Len of moves is %d.\n", len(p.data.moves))

	if len(p.data.moves) == 1 {
		fmt.Printf("Added second move.\n")
		moves = append(moves, &Square{x: start.x, y: start.y + 2})
	}

	// TODO: when we add capturing, just chop this and validate the moves
	// one by one, since the capture rules are weird.

	validMoves := make([]*Square, 0)
	for _, move := range moves {
		if !move.isValid() {
			continue
		}
		status := p.pieceData().evaluateSquare(move)
		if status == squareVacant || status == squareOccupiedByOpponent {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}

func (p Pawn) getShorthand() string {
	return "p"
}
