package main

import (
	"errors"
	"fmt"
)

type Board struct {
	// Why + 1?  Chess squares are 1-indexed, for some dumb reason.
	// We embrace that 1-offset rather than subtracting everywhere.
	squares [endSquare + 1][endSquare + 1]*Piece

	// Easy access to kings, indexed by color.
	kings [2]*Piece

	// Unordered array of pieces, indexed by color.
	colorPieces [2][numPiecesPerPlayer]*Piece
	players     [2]string

	moveCount int
}

// Instantiates a new board.
func NewBoard() *Board {
	b := new(Board)
	b.populatePieces(White)
	b.populatePieces(Black)
	b.moveCount = 0

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
		b.kings[piece.color] = piece
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
	fmt.Printf("\nWhite player: %s, Black player: %s\n", b.players[White], b.players[Black])

	fmt.Printf("   a   b   c   d   e   f   g   h\n")
	for y := endSquare; y >= startSquare; y-- {
		fmt.Printf("%d ", y)
		for x := startSquare; x <= endSquare; x++ {
			p := b.getPieceByCoordinates(x, y)
			if p != nil {
				if p.color == White {
					fmt.Printf(" w%s ", string(p.pieceType))
				} else if p.color == Black {
					fmt.Printf(" b%s ", string(p.pieceType))
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
	colors := []Color{White, Black}
	for _, color := range colors {
		if b.isKingInCheck(color) {
			// If any piece has a valid move, then color is not mated.
			for _, piece := range b.colorPieces[color] {
				possibleMoves := piece.generateMoves(*piece.getSquare())
				if len(possibleMoves) > 0 {
					if color == White {
						return WhiteInCheck
					} else {
						return BlackInCheck
					}
				}
			}
			if color == White {
				return WhiteCheckmated
			} else {
				return BlackCheckmated
			}
		}
	}

	// TODO: What about a draw?  Think there's an edge case here.
	return GameOn
}

// Finds the king for a color.
func (b Board) getKing(color Color) *Piece {
	return b.kings[color]
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
	pieceArr := &b.colorPieces[color]
	if color == Black {
		pawnRow = endSquare - 1
		rookRow = endSquare
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

// Allows you to set a player name.  Used for playing actual games.
func (b *Board) setPlayerName(color Color, name string) {
	b.players[color] = name
}

// Which player moves next?  Used for actual games.
func (b Board) getPlayerWithNextMove() string {
	return b.players[b.getColorWithNextMove()]
}

// Which color moves next?
func (b Board) getColorWithNextMove() Color {
	if b.moveCount%2 == 0 {
		return White
	}
	return Black
}

// Attempt to move a piece to a user-provided location.
func (b Board) attemptUserMove(move string) (bool, error) {
	// TODO: Maybe this belongs in gopherchess.go?  Not really related to board.
	minMoveLength := 3
	if len(move) < minMoveLength {
		return false, errors.New("Move must be at least 3 characters.")
	}

	// First, validate piece type.
	notationPiece := move[0]
	pieceType, err := getPieceTypeFromNotation(notationPiece)
	if err != nil {
		return false, errors.New("Notation specifies an invalid piece type.")
	}

	// Then, validate the destination square.
	notationSquare := move[1:3]
	square, err := getSquareFromNotation(notationSquare)
	if err != nil {
		return false, errors.New("Notation specifies an invalid square.")
	}

	// Next, find the piece and that can move to the destination square.
	movingColor := b.getColorWithNextMove()
	movingColorPieces := b.colorPieces[movingColor]
	for _, p := range movingColorPieces {
		if p.pieceType != pieceType {
			continue
		}

		// We have found the square! Let's make a move.
		moveErr := p.move(square)
		if moveErr == nil {
			return true, nil
		}
	}

	return false, errors.New("No piece found who can move to specified square.")
}
