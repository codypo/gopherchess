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

func (b Board) evaluateSquare(c Color, s *Square) SquareState {
	if !s.isValid() {
		return SquareInvalid
	}

	// Does either side have a piece on this square?
	pieceOnSquare := b.getPieceBySquare(s.x, s.y)
	if pieceOnSquare != nil {
		if pieceOnSquare.color == c {
			return SquareOccupiedByMe
		} else {
			// Can do check check. HAR HAR HAR.
			return SquareOccupiedByOpponent
		}
	}

	return SquareVacant
}

func (b Board) getPieceBySquare(x int, y int) *Piece {
	// Does either side have a piece on this square?
	if bPiece, _ := b.getPlayer(Black).getPieceByCoordinate(x, y); bPiece != nil {
		return bPiece
	}

	if wPiece, _ := b.getPlayer(White).getPieceByCoordinate(x, y); wPiece != nil {
		return wPiece
	}

	return nil
}

// Print out the current state of the board.  Useful in the event
// this thing can ever play a game.
func (b Board) prettyPrint() {
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// TODO: should also print out row and column designations,
	// and distinguish pieces by color.
	for y := startSquare; y <= endSquare; y++ {
		for x := startSquare; x <= endSquare; x++ {
			if wp, _ := white.getPieceByCoordinate(x, y); wp != nil {
				fmt.Printf(" %s ", wp.getShorthand())
			} else if bp, _ := black.getPieceByCoordinate(x, y); bp != nil {
				fmt.Printf(" %s ", bp.getShorthand())
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Printf("\n")
	}
}

// Returns the state the board is in.  Interesting returns here are
// check or check mate.
func (b Board) getGameState() GameState {
	return Default
}
