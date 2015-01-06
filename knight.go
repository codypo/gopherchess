package main

import "fmt"

type Knight struct {
	data *PieceData
}

func (k Knight) move(newSquare Square) bool {
	return false
}

func (k Knight) pieceData() *PieceData {
	return k.data
}

func (k Knight) generateValidMoves(start Square) []*Square {
	moves := make([]*Square, 0)

	move := &Square{x: start.x + 2, y: start.y + 1}
	moves = append(moves, move)
	move = &Square{x: start.x + 2, y: start.y - 1}
	moves = append(moves, move)
	move = &Square{x: start.x - 2, y: start.y + 1}
	moves = append(moves, move)
	move = &Square{x: start.x - 2, y: start.y - 1}
	moves = append(moves, move)

	move = &Square{x: start.x + 1, y: start.y + 2}
	moves = append(moves, move)
	move = &Square{x: start.x + 1, y: start.y - 2}
	moves = append(moves, move)
	move = &Square{x: start.x - 1, y: start.y + 2}
	moves = append(moves, move)
	move = &Square{x: start.x - 1, y: start.y - 2}
	moves = append(moves, move)

	validMoves := make([]*Square, 0)

	for _, move := range moves {
		status := k.pieceData().evaluateSquare(move)
		if status == squareVacant || status == squareOccupiedByOpponent {
			validMoves = append(validMoves, move)
			fmt.Printf("Knight - Adding valid move %d, %d\n", move.x, move.y)
		}
	}

	return validMoves
}
