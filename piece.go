package main

import (
	"fmt"
)

// The ability to generate moves (e.g., a bishop's moves vs. a
// knight's moves) is really the only difference between pieces.
// Encapsulate this in an interface that each piece type can
// implement.
type PieceMover interface {
	generateMoves(start Square) []*Square
	getShorthand() string
	getPiece() *Piece
	setPiece(piece *Piece)
}

// Piece struct encapsulates all of the common data that's
// known about a piece.  Color, location, past moves, etc.
type Piece struct {
	color     Color
	captured  bool
	moves     []*Square
	board     *Board
	mover     PieceMover
	pieceType PieceType
}

// Gets a piece's current square.  If the piece is captured,
// returns nil.
func (p *Piece) getSquare() *Square {
	if p.captured {
		return nil
	}
	return p.moves[len(p.moves)-1]
}

// You been captured, piece!
func (p *Piece) setCaptured() {
	p.captured = true
}

// Moves a piece to a square, if allowed.  If that square is
// occupied by the opponent, thsi will capture the occupant.
func (p *Piece) move(square *Square) error {
	// First, confirm this is a valid move.
	if !p.canMoveToSquare(*square) {
		return fmt.Errorf("Specified move is not valid.")
	}

	// TODO: this is inefficient.  generateMoves typically
	// calls evaluateSquare.  generateMoves could instead
	// return a hash map of square to status.
	moveStatus := p.board.evaluateSquare(p.color, square)
	switch moveStatus {
	case SquareOccupiedByOpponent:
		capturedPiece := p.board.getPieceBySquare(*square)
		capturedPiece.setCaptured()
	case SquareVacant:
		p.moves = append(p.moves, square)
		break
	default:
		return fmt.Errorf("Specified move is not valid.")
	}

	return nil
}

// Used only for unit tests.
func (p *Piece) forceMove(square *Square) {
	p.moves = append(p.moves, square)
}

// Does a piece's current square match given coordinates.
func (p Piece) matchesCoordinates(x int, y int) bool {
	return (p.getSquare().x == x) && (p.getSquare().y == y)
}

// Utility func to return a piece's y coordinate.
func (p Piece) y() int {
	return p.getSquare().y
}

// Utility func to return a piece's x coordinate.
func (p Piece) x() int {
	return p.getSquare().x
}

// Determines the state of a square, from the perspective
// of this piece.
func (p Piece) evaluateSquare(square *Square) SquareState {
	return p.board.evaluateSquare(p.color, square)
}

// Determines if this piece can move to a given square.
func (p Piece) canMoveToSquare(square Square) bool {
	for _, s := range p.generateMoves(*p.getSquare()) {
		if s.x == square.x && s.y == square.y {
			return true
		}
	}

	return false
}

// Determines if moving piece to specified square results in
// my side getting checked.  That's not cool.
func (p Piece) doesMoveEndangerKing(square Square) bool {
	// Create a copy of the board where we've already made
	// this move.  Evaluate if my color is in check.
	boardClone := p.board.deepCopy()
	clonePiece := boardClone.getPieceBySquare(*p.getSquare())
	clonePiece.forceMove(&square)
	return boardClone.doesMoveLeadToCheck(clonePiece)

	// RECURSION!!!!
	// doesMoveEndangerKing calls getGameState - must make sure we're not moving into check
	// getGameState calls player.canMoveToSquare - must make sure oppo player can take king from new spot
	// player.canMoveToSquare calls piece.canMoveToSquare - must make see if any piece can take king with this change.  Can I exclude moving piece?
	// piece.canMoveToSquare calls generateMoves - must see if my piece can move to this square.
	// generateMoves calls doesMoveEndangerKing - must see if I'm exposing king.
}

// Gets the shorthand notation for a piece, like p for Pawn.
func (p Piece) getShorthand() string {
	return p.mover.getShorthand()
}

// Generate all of the valid moves for a piece, given its
// starting square.
func (p Piece) generateMoves(start Square) []*Square {
	// Go through our list of generated moves.  Verify that
	// none of these moves endangers our king.
	moves := make([]*Square, 0)
	for _, move := range p.mover.generateMoves(*p.getSquare()) {
		if !p.doesMoveEndangerKing(*move) {
			moves = append(moves, move)
		}
	}
	return moves
}

// Generate all valid moves on the board for a given piece in only
// the diagonal directions.
// Reusable across pieces (bishop, queen), so it lives here.
func (p Piece) generateDiagonalMoves(start Square) []*Square {
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
			if status == SquareVacant || status == SquareOccupiedByOpponent {
				moves = append(moves, move)
			}

			// You can only go onward if you're looking at a vacant square.
			goUpLeft = (status == SquareVacant)
		}

		// Evaluate the next move down and to the left.
		if goDownLeft {
			move := &Square{x: start.x - i, y: start.y - i}
			status := p.evaluateSquare(move)
			if status == SquareVacant || status == SquareOccupiedByOpponent {
				moves = append(moves, move)
			}

			goDownLeft = (status == SquareVacant)
		}

		// Evaluate the next move up and to the right.
		if goUpRight {
			move := &Square{x: start.x + i, y: start.y + i}
			status := p.evaluateSquare(move)
			if status == SquareVacant || status == SquareOccupiedByOpponent {
				moves = append(moves, move)
			}
			goUpRight = (status == SquareVacant)
		}

		// Evaluate the next move down and to the right
		if goDownRight {
			move := &Square{x: start.x + i, y: start.y - i}
			status := p.evaluateSquare(move)
			if status == SquareVacant || status == SquareOccupiedByOpponent {
				moves = append(moves, move)
			}
			goDownRight = (status == SquareVacant)
		}
	}

	return moves
}

// Generate all valid moves on the board for a given piece in only
// the vertical and horizontal directions directions.
// Reusable across pieces (queen, rook), so it lives here.
func (p Piece) generateStraightMoves(start Square) []*Square {
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
			if status == SquareVacant || status == SquareOccupiedByOpponent {
				moves = append(moves, move)
			}

			// You can only go onward if you're looking at a vacant square.
			goUp = (status == SquareVacant)
		}

		// Evaluate the next move down.
		if goDown {
			move := &Square{x: start.x, y: start.y - i}
			status := p.evaluateSquare(move)
			if status == SquareVacant || status == SquareOccupiedByOpponent {
				moves = append(moves, move)
			}

			goDown = (status == SquareVacant)
		}

		// Evaluate the next move right.
		if goRight {
			move := &Square{x: start.x + i, y: start.y}
			status := p.evaluateSquare(move)
			if status == SquareVacant || status == SquareOccupiedByOpponent {
				moves = append(moves, move)
			}
			goRight = (status == SquareVacant)
		}

		// Evaluate the next move left.
		if goLeft {
			move := &Square{x: start.x - i, y: start.y}
			status := p.evaluateSquare(move)
			if status == SquareVacant || status == SquareOccupiedByOpponent {
				moves = append(moves, move)
			}
			goLeft = (status == SquareVacant)
		}
	}

	return moves
}

func (p Piece) deepCopy(board *Board) *Piece {
	c := new(Piece)
	c.color = p.color
	c.captured = p.captured
	c.board = board
	c.pieceType = p.pieceType
	c.mover = p.mover
	c.mover.setPiece(c)

	// Note that we only copy the square's current position.
	// Not sure if this is needed.
	c.moves = []*Square{p.getSquare()}

	return c
}

func NewPiece(color Color, square *Square, board *Board, pieceType PieceType) *Piece {
	p := new(Piece)
	p.color = color
	p.captured = false
	p.moves = []*Square{square}
	p.board = board

	p.pieceType = pieceType
	switch pieceType {
	case PawnType:
		p.mover = new(Pawn)
	case RookType:
		p.mover = new(Rook)
	case BishopType:
		p.mover = new(Bishop)
	case KnightType:
		p.mover = new(Knight)
	case QueenType:
		p.mover = new(Queen)
	case KingType:
		p.mover = new(King)

	}
	p.mover.setPiece(p)

	//fmt.Printf(" 3 my piece is %s ???\n", p)

	return p
}
