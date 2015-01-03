package main

import "fmt"

type Rook struct {
	data *PieceData
}

func (r Rook) move(newSquare Square) bool {
	return false
}

func (r Rook) pieceData() *PieceData {
	return r.data
}

func (r Rook) generateValidMoves(start Square) []*Square {
	fmt.Printf("I am in %d, %d...\n", start.x, start.y)
	moves := make([]*Square, 0)

	proceedUp := true
	proceedDown := true
	proceedLeft := true
	proceedRight := true

	for i := startSquare; i <= endSquare; i++ {
		// Evaluate the next move up.
		if proceedUp {
			move := &Square{x: start.x, y: start.y + i}
			status := r.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
				fmt.Printf("Adding valid move %d, %d\n", move.x, move.y)
			} else {
				proceedUp = false
			}
		}

		// Evaluate the next move down.
		if proceedDown {
			move := &Square{x: start.x, y: start.y - i}
			status := r.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
				fmt.Printf("Adding valid move %d, %d\n", move.x, move.y)
			} else {
				proceedDown = false
			}
		}

		// Evaluate the next move right.
		if proceedRight {
			move := &Square{x: start.x + i, y: start.y}
			status := r.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
				fmt.Printf("Adding valid move %d, %d\n", move.x, move.y)
			} else {
				proceedRight = false
			}
		}

		// Evaluate the next move left.
		if proceedLeft {
			move := &Square{x: start.x - i, y: start.y}
			status := r.pieceData().evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
				fmt.Printf("Adding valid move %d, %d\n", move.x, move.y)
			} else {
				proceedLeft = false
			}
		}

		fmt.Printf("Length of moves is %d\n", len(moves))
	}

	return moves
}
