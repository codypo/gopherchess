package main

import "fmt"

type Board struct {
	// Why + 1?  Chess squares are 1-indexed, for some dumb reason.
	// We embrace that 1-offset rather than subtracting everywhere.
	squares     [endSquare + 1][endSquare + 1]*Piece
	whiteKing   *Piece
	blackKing   *Piece
	whitePieces [numPiecesPerPlayer]*Piece
	blackPieces [numPiecesPerPlayer]*Piece
}

// Instantiates a new board.
func NewBoard() *Board {
	b := new(Board)
	b.populatePieces(White)
	b.populatePieces(Black)

	return b
}

// Move the specified piece to the appropriate location in the appropriate
// square space.
func (b *Board) updateSquare(piece *Piece) {
	// First, nil out the square this dude was formerly occupying.
	if len(piece.moves) > 1 {
		oldSquare := piece.moves[len(piece.moves)-2]
		b.squares[oldSquare.x][oldSquare.y] = nil
	}

	// Then, update the new square.
	b.squares[piece.x()][piece.y()] = piece

	// If this happens to be a king, update its pointer.
	if piece.pieceType == KingType {
		if piece.color == White {
			b.whiteKing = piece
		} else {
			b.blackKing = piece
		}
	}
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
	return b.squares[x][y]
}

// Given a square, fetch its occupying piece.
func (b Board) getPieceBySquare(s Square) *Piece {
	return b.getPieceByCoordinates(s.x, s.y)
}

// Print out the current state of the board.  Useful in the event
// this thing can ever play a game.
func (b Board) prettyPrint() {
	// TODO: should also print out row and column designations,
	// and distinguish pieces by color.
	for y := endSquare; y >= startSquare; y-- {
		for x := startSquare; x <= endSquare; x++ {
			p := b.getPieceByCoordinates(x, y)
			if p != nil {
				if p.color == White {
					fmt.Printf(" w%s ", p.getShorthand())
				} else if p.color == Black {
					fmt.Printf(" b%s ", p.getShorthand())
				}
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
	// TODO: Collapse this method.
	if b.isKingInCheck(White) {
		// If any piece has a valid move, then white is not mated.
		for _, wPiece := range b.whitePieces {
			possibleMoves := wPiece.generateMoves(*wPiece.getSquare())
			if len(possibleMoves) > 0 {
				return WhiteInCheck
			}
		}
		return WhiteCheckmated
	}
	if b.isKingInCheck(Black) {
		// If any piece has a valid move, then white is not mated.
		for _, bPiece := range b.blackPieces {
			possibleMoves := bPiece.generateMoves(*bPiece.getSquare())
			if len(possibleMoves) > 0 {
				return BlackInCheck
			}
		}
		return BlackCheckmated
	}

	// TODO: What about a draw?  Think there's an edge case here.
	return GameOn
}

// Finds the king for a color.
func (b Board) getKing(color Color) *Piece {
	if color == White {
		return b.whiteKing
	}
	return b.blackKing
}

// Is the king now in check?  Use the king's position to evaluate the
// squares around it and see if attacking pieces are there.
func (b Board) isKingInCheck(color Color) bool {
	// Opposing pawn is descending our y axis.
	oppoPawnDirection := -1
	if color == Black {
		oppoPawnDirection = 1
	}

	myKing := b.getKing(color)

	// From the king's position, generate diagonal moves.  Do we
	// see an opposing bishop, pawn, or queen?
	diagMoves := myKing.generateDiagonalMoves(*myKing.getSquare())
	for _, m := range diagMoves {
		p := b.getPieceBySquare(*m)
		if p == nil || p.color == color {
			continue
		}
		if p.pieceType == BishopType || p.pieceType == QueenType {
			return true
		}
		if p.pieceType == PawnType {
			// Opposing pawns can only go in 1 direction.
			if (myKing.getSquare().y - m.y) == oppoPawnDirection {
				return true
			}
		}
	}

	// From the king's position, generate straight moves.  Do we
	// see an opposing rook or queen?
	straightMoves := myKing.generateStraightMoves(*myKing.getSquare())
	for _, m := range straightMoves {
		p := b.getPieceBySquare(*m)
		if p == nil || p.color == color {
			continue
		}
		if p.pieceType == RookType || p.pieceType == QueenType {
			return true
		}
	}

	// From the king's position, generate knight moves.  Do we see
	// an opposing knight?
	knightMoves := myKing.generateKnightMoves(*myKing.getSquare())
	for _, m := range knightMoves {
		p := b.getPieceBySquare(*m)
		if p == nil || p.color == color {
			continue
		}
		if p.pieceType == KnightType {
			return true
		}
	}
	return false
}

// Does a deep copy of the board.  Used to assess hypothetical moves.
func (b Board) deepCopy() Board {
	c := new(Board)
	for x := startSquare; x <= endSquare; x++ {
		for y := startSquare; y <= endSquare; y++ {
			if b.squares[x][y] == nil {
				continue
			}
			c.squares[x][y] = b.squares[x][y].deepCopy(c)
		}
	}

	return *c
}

func (b Board) dumpSquares() {
	for x := startSquare; x <= endSquare; x++ {
		for y := startSquare; y <= endSquare; y++ {
			p := b.squares[x][y]
			if p == nil {
				fmt.Printf("%d, %d is nil\n", x, y)
			} else {
				fmt.Printf("DS %d, %d is %T: &i=%p IIII=%v\n", p.x(), p.y(), p, &p, p)
			}
		}
	}
}

// Populate all of the pieces for a color.
func (b *Board) populatePieces(color Color) {
	pieceIndex := 0

	// Note that NewPiece handles the positioning of the piece
	// on the board, which populates the squares arrays.

	// p p p p p p p p.
	// r k b q k b k r.
	pawnRow := startSquare + 1
	rookRow := startSquare
	pieceArr := &b.whitePieces
	if color == Black {
		pawnRow = endSquare - 1
		rookRow = endSquare
		pieceArr = &b.blackPieces
	}

	s := &Square{x: startSquare, y: startSquare}

	// Populate pawns.
	for x := startSquare; x <= endSquare; x++ {
		s = &Square{x: x, y: pawnRow}
		NewPiece(color, s, b, PawnType)
		pieceIndex++
	}

	// Populate rooks.
	s = &Square{x: startSquare, y: rookRow}
	NewPiece(color, s, b, RookType)
	pieceIndex++

	// TODO: This pieceIndex part is silly.
	s = &Square{x: endSquare, y: rookRow}
	NewPiece(color, s, b, RookType)
	pieceIndex++

	// Populate knights.
	s = &Square{x: startSquare + 1, y: rookRow}
	NewPiece(color, s, b, KnightType)
	pieceIndex++

	s = &Square{x: endSquare - 1, y: rookRow}
	NewPiece(color, s, b, KnightType)
	pieceIndex++

	// Populate bishops.
	s = &Square{x: startSquare + 2, y: rookRow}
	NewPiece(color, s, b, BishopType)
	pieceIndex++

	s = &Square{x: endSquare - 2, y: rookRow}
	NewPiece(color, s, b, BishopType)
	pieceIndex++

	// Populate the queen.
	s = &Square{x: startSquare + 3, y: rookRow}
	NewPiece(color, s, b, QueenType)
	pieceIndex++

	// Populate the king.
	s = &Square{x: startSquare + 4, y: rookRow}
	NewPiece(color, s, b, KingType)
	pieceIndex++

	// Populate the arrays used to track each color's pieces.
	for i := startSquare; i <= endSquare; i++ {
		// Ugh, stupid 1 offset.  We want a zero indexed array here.
		noOffset := i - 1
		pieceArr[noOffset] = b.squares[i][rookRow]
		pieceArr[noOffset+endSquare] = b.squares[i][pawnRow]
	}
}
