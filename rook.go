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

	// HACK: this is not at all efficient.
	immediateMoves := make([]*Square, 4)
	immediateMoves[0] = &Square{x: start.x - 1, y: start.y}
	immediateMoves[1] = &Square{x: start.x + 1, y: start.y}
	immediateMoves[2] = &Square{x: start.x, y: start.y + 1}
	immediateMoves[3] = &Square{x: start.x, y: start.y - 1}
	immediateOpenings := 0

	for _, i := range immediateMoves {
		status := r.pieceData().evaluateSquare(i)
		if status == squareVacant {
			immediateOpenings++
		} else if status == squareOccupiedByOpponent {
			immediateOpenings++
		}
	}

	fmt.Printf("immediateOpenings is of length %d\n", immediateOpenings)
	if immediateOpenings == 0 {
		return moves
	}

	possibleMoves := [endSquare * 2]*Square{}
	pIndex := 0

	// Rooks can move horizontally or vertically.
	// Look at the possible horizontal moves.
	for x := startSquare; x <= endSquare; x++ {
		possibleMoves[pIndex] = &Square{x: x, y: start.y}
		pIndex++
	}

	// Look at the possible vertical moves.
	for y := startSquare; y <= endSquare; y++ {
		possibleMoves[pIndex] = &Square{x: start.x, y: y}
		pIndex++
	}

	// Can the rook actually move to any of these places?
	for _, s := range possibleMoves {
		fmt.Printf("Evaluating possible move %d, %d\n", s.x, s.y)
		status := r.pieceData().evaluateSquare(s)
		if status == squareVacant {
			fmt.Printf(" Appending vacant square\n")
			moves = append(moves, s)
		} else if status == squareOccupiedByMe {
			// Pass.  Castle at some point?
		} else if status == squareOccupiedByOpponent {
			fmt.Printf(" Appending occupied square\n")
			moves = append(moves, s)
		}
	}

	// Still need to catch the fact that rooks can't hop over dudes.
	fmt.Printf("Length of moves is %d, possibleMoves is %d\n", len(moves), len(possibleMoves))
	return moves
}
