package main

import "fmt"

type Piece interface {
	pieceData() *PieceData
	generateMoves(start Square) []*Square
	getShorthand() string
}

// Piece data exists because there's no concept of an abstract
// class in go.  So, we have the piece interface and the piece
// data struct.
type PieceData struct {
	color    Color
	captured bool
	moves    []*Square
	board    *Board
}

func (p PieceData) getSquare() *Square {
	return p.moves[len(p.moves)-1]
}

func (p PieceData) move(square *Square) {
	fmt.Printf("\nAdding... Length of moves is %d\n", len(p.moves))
	p.moves = append(p.moves, square)
	fmt.Printf("Added!  Length of moves is now %d\n\n", len(p.moves))
}

func (p PieceData) matchesCoordinates(x int, y int) bool {
	return (p.getSquare().x == x) && (p.getSquare().y == y)
}

func (p PieceData) y() int {
	return p.getSquare().y
}

func (p PieceData) x() int {
	return p.getSquare().x
}

func (p PieceData) evaluateSquare(square *Square) int {
	return p.board.evaluateSquare(p.color, square)
}

// Generate all valid moves on the board for a given piece in only
// the diagonal directions.
func (p PieceData) generateDiagonalMoves(start Square) []*Square {
	moves := make([]*Square, 0)

	goUpLeft := true
	goDownLeft := true
	goUpRight := true
	goDownRight := true

	for i := startSquare; i <= endSquare; i++ {
		// Evaluate the next move up and to the left.
		if goUpLeft {
			move := &Square{x: start.x - i, y: start.y + i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			// You can only go onward if you're looking at a vacant square.
			goUpLeft = (status == squareVacant)
		}

		// Evaluate the next move down and to the left.
		if goDownLeft {
			move := &Square{x: start.x - i, y: start.y - i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			goDownLeft = (status == squareVacant)
		}

		// Evaluate the next move up and to the right.
		if goUpRight {
			move := &Square{x: start.x + i, y: start.y + i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			goUpRight = (status == squareVacant)
		}

		// Evaluate the next move down and to the right
		if goDownRight {
			move := &Square{x: start.x + i, y: start.y - i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			goDownRight = (status == squareVacant)
		}
	}

	return moves
}

// Generate all valid moves on the board for a given piece in only
// the verticl and horizontal directions directions.
func (p PieceData) generateStraightMoves(start Square) []*Square {
	moves := make([]*Square, 0)

	goUp := true
	goDown := true
	goLeft := true
	goRight := true

	for i := startSquare; i <= endSquare; i++ {
		// Evaluate the next move up.
		if goUp {
			move := &Square{x: start.x, y: start.y + i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			// You can only go onward if you're looking at a vacant square.
			goUp = (status == squareVacant)
		}

		// Evaluate the next move down.
		if goDown {
			move := &Square{x: start.x, y: start.y - i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			goDown = (status == squareVacant)
		}

		// Evaluate the next move right.
		if goRight {
			move := &Square{x: start.x + i, y: start.y}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			goRight = (status == squareVacant)
		}

		// Evaluate the next move left.
		if goLeft {
			move := &Square{x: start.x - i, y: start.y}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			goLeft = (status == squareVacant)
		}
	}

	return moves
}

func NewPieceData(color Color, square *Square, board *Board) *PieceData {
	pd := new(PieceData)
	pd.color = color
	pd.captured = false
	pd.moves = []*Square{square}
	pd.board = board
	return pd
}
