package main

type Piece interface {
	move(newSquare Square) bool
	pieceData() *PieceData
	generateValidMoves(start Square) []*Square
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

func (p PieceData) addMove(square *Square) {
	p.moves = append(p.moves, square)
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

func (p PieceData) generateDiagonalMoves(start Square) []*Square {
	moves := make([]*Square, 0)

	proceedUpLeft := true
	proceedDownLeft := true
	proceedUpRight := true
	proceedDownRight := true

	for i := startSquare; i <= endSquare; i++ {
		// Evaluate the next move up and to the left.
		if proceedUpLeft {
			move := &Square{x: start.x - i, y: start.y + i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			// You can only proceed onward if you're looking at a vacant square.
			proceedUpLeft = (status == squareVacant)
		}

		// Evaluate the next move down and to the left.
		if proceedDownLeft {
			move := &Square{x: start.x - i, y: start.y - i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			proceedDownLeft = (status == squareVacant)
		}

		// Evaluate the next move up and to the right.
		if proceedUpRight {
			move := &Square{x: start.x + i, y: start.y + i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedUpRight = (status == squareVacant)
		}

		// Evaluate the next move down and to the right
		if proceedDownRight {
			move := &Square{x: start.x + i, y: start.y - i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedDownRight = (status == squareVacant)
		}
	}

	return moves
}

func (p PieceData) generateStraightMoves(start Square) []*Square {
	moves := make([]*Square, 0)

	proceedUp := true
	proceedDown := true
	proceedLeft := true
	proceedRight := true

	for i := startSquare; i <= endSquare; i++ {
		// Evaluate the next move up.
		if proceedUp {
			move := &Square{x: start.x, y: start.y + i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			// You can only proceed onward if you're looking at a vacant square.
			proceedUp = (status == squareVacant)
		}

		// Evaluate the next move down.
		if proceedDown {
			move := &Square{x: start.x, y: start.y - i}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}

			proceedDown = (status == squareVacant)
		}

		// Evaluate the next move right.
		if proceedRight {
			move := &Square{x: start.x + i, y: start.y}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedRight = (status == squareVacant)
		}

		// Evaluate the next move left.
		if proceedLeft {
			move := &Square{x: start.x - i, y: start.y}
			status := p.evaluateSquare(move)
			if status == squareVacant || status == squareOccupiedByOpponent {
				moves = append(moves, move)
			}
			proceedLeft = (status == squareVacant)
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
