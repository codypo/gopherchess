package main

import "fmt"

type Board struct {
	players [numPlayers]Player
	squares [endSquare][endSquare]Square
}

func NewBoard() *Board {
	b := new(Board)

	// This feels gross, but it's a way to ensure that 0-indexed array
	// lines up with silly 1-indexed squares.
	// TODO: We could actually ditch this.  We don't use these squares.
	offset := startSquare - 0
	for x := startSquare; x <= endSquare; x++ {
		for y := startSquare; y <= endSquare; y++ {
			s := Square{x: x, y: y}
			b.squares[x-offset][y-offset] = s
		}
	}

	p0 := NewPlayer(White, b)
	b.players[0] = *p0

	p1 := NewPlayer(Black, b)
	b.players[1] = *p1

	return b
}

func (b Board) getPlayer(color Color) Player {
	for _, p := range b.players {
		if p.color == color {
			return p
		}
	}

	// HAAAAAAAAAAAACK
	return b.players[0]
}

// Determines the state of a square.  We must know the color
// of the moving side, so we can determine if the square is
// occupied.
func (b Board) evaluateSquare(c Color, s *Square) SquareState {
	if !s.isValid() {
		return SquareInvalid
	}

	// Does either side have a piece on this square?
	if piece := b.getPieceBySquare(*s); piece != nil {
		if piece.color == c {
			return SquareOccupiedByMe
		} else {
			// Can do check check. HAR HAR HAR.
			return SquareOccupiedByOpponent
		}
	}

	return SquareVacant
}

// Given the coordinates of a square, fetch its occuping piece.
func (b Board) getPieceByCoordinates(x int, y int) *Piece {
	// Does either side have a piece on this square?
	if bPiece, _ := b.getPlayer(Black).getPieceByCoordinate(x, y); bPiece != nil {
		return bPiece
	}

	if wPiece, _ := b.getPlayer(White).getPieceByCoordinate(x, y); wPiece != nil {
		return wPiece
	}

	return nil
}

// Given a square, fetch its occupying piece.
func (b Board) getPieceBySquare(s Square) *Piece {
	return b.getPieceByCoordinates(s.x, s.y)
}

// Print out the current state of the board.  Useful in the event
// this thing can ever play a game.
func (b Board) prettyPrint() {
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// TODO: should also print out row and column designations,
	// and distinguish pieces by color.
	for y := endSquare; y >= startSquare; y-- {
		for x := startSquare; x <= endSquare; x++ {
			if wp, _ := white.getPieceByCoordinate(x, y); wp != nil {
				fmt.Printf(" w%s ", wp.getShorthand())
			} else if bp, _ := black.getPieceByCoordinate(x, y); bp != nil {
				fmt.Printf(" b%s ", bp.getShorthand())
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Printf("\n")
	}
}

// Returns the state the board is in.  Interesting returns here are
// check or check mate.
func (b Board) getGameState() GameState {
	if b.isKingInCheck(White) {
		return WhiteInCheck
	}
	if b.isKingInCheck(Black) {
		return BlackInCheck
	}

	return GameOn
}

// Is the king now in check?  Use the king's position to evaluate the
// squares around it and see if attacking pieces are there.
func (b Board) isKingInCheck(color Color) bool {
	// Opposing pawn is descending our y axis.
	oppoPawnDirection := -1

	// TODO: This is dumb.  Just use color.
	myPlayer := b.players[0]
	if color == Black {
		myPlayer = b.players[1]
		oppoPawnDirection = 1
	}

	myKing, _ := myPlayer.getKing()

	// From the king's position, generate diagonal moves.  Do we
	// see an opposing bishop, pawn, or queen?
	diagMoves := myKing.generateDiagonalMoves(*myKing.getSquare())
	for _, m := range diagMoves {
		p := b.getPieceBySquare(*m)
		if p == nil {
			continue
		}
		if p.pieceType == BishopType || p.pieceType == QueenType {
			return true
		}
		if p.pieceType == PawnType {
			// Opposing pawns can only go in 1 direction.
			if m.y-myKing.getSquare().y == oppoPawnDirection {
				return true
			}
		}
	}

	// From the king's position, generate straight moves.  Do we
	// see an opposing rook or queen?
	straightMoves := myKing.generateStraightMoves(*myKing.getSquare())
	for _, m := range straightMoves {
		p := b.getPieceBySquare(*m)
		if p == nil {
			continue
		}
		if p.pieceType == RookType || p.pieceType == QueenType {
			return true
		}
	}

	// TODO: What about those wacky knights?
	return false
}

// Does a deep copy of the board.  Used to assess hypothetical moves.
func (b Board) deepCopy() Board {
	c := new(Board)
	c.players[0] = *b.players[0].deepCopy(c, White)
	c.players[1] = *b.players[1].deepCopy(c, Black)

	return *c
}
