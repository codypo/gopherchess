package main

import (
	"testing"
)

func TestBoardNew(t *testing.T) {
	b := NewBoard()
	totalSquares := 0
	for i, _ := range b.squares {
		totalSquares += len(b.squares[i])
	}
	// TODO: This is weird.  We build a 9 by 9 board so we can index
	// directly into our arrays of chess squares, rather than doing
	// a 1 offset.  Not sure if I like that.
	if totalSquares != 81 {
		t.Errorf("Board contains %d squares.", totalSquares)
	}
}

func TestBoardRecognizesWhiteInCheckByRook(t *testing.T) {
	b := NewBoard()

	// Force the white king to the middle of the board.
	wKing := b.getKing(White)
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where white king is in check.
	bRook := b.getPieceByCoordinates(8, 8)
	bRook.forceMove(&Square{x: 8, y: 4})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByRook(t *testing.T) {
	b := NewBoard()

	// Force the king to the middle of the board.
	bKing := b.getKing(Black)
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where king is in check.
	wRook := b.getPieceByCoordinates(1, 1)
	wRook.forceMove(&Square{x: 1, y: 4})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardRecognizesWhiteInCheckByQueen(t *testing.T) {
	b := NewBoard()

	// Force the white king to the middle of the board.
	wKing := b.getKing(White)
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force queen into a spot where white king is in check.
	bQueen := b.getPieceByCoordinates(4, 8)
	bQueen.forceMove(&Square{x: 6, y: 6})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByQueen(t *testing.T) {
	b := NewBoard()

	// Force the king to the middle of the board.
	bKing := b.getKing(Black)
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where king is in check.
	wQueen := b.getPieceByCoordinates(4, 1)
	wQueen.forceMove(&Square{x: 3, y: 3})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardRecognizesWhiteInCheckByBishop(t *testing.T) {
	b := NewBoard()

	// Force the white king to the middle of the board.
	wKing := b.getKing(White)
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force bishop into a spot where white king is in check.
	bBishop := b.getPieceByCoordinates(3, 8)
	bBishop.forceMove(&Square{x: 6, y: 6})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByBishop(t *testing.T) {
	b := NewBoard()

	// Force the king to the middle of the board.
	bKing := b.getKing(Black)
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force bishop into a spot where king is in check.
	wBishop := b.getPieceByCoordinates(6, 1)
	wBishop.forceMove(&Square{x: 5, y: 5})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardRecognizesWhiteInCheckByPawn(t *testing.T) {
	b := NewBoard()

	// Force the white king to the middle of the board.
	wKing := b.getKing(White)
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force pawn into a spot where white king is in check.
	bPawn := b.getPieceByCoordinates(3, 7)
	bPawn.forceMove(&Square{x: 3, y: 5})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByPawn(t *testing.T) {
	b := NewBoard()

	// Force the king to the middle of the board.
	bKing := b.getKing(Black)
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force pawn into a spot where king is in check.
	wPawn := b.getPieceByCoordinates(6, 2)
	wPawn.forceMove(&Square{x: 3, y: 3})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardDoesNotRecognizeWhiteInCheckByBackwardsPawn(t *testing.T) {
	b := NewBoard()

	// Force the white king to the middle of the board.
	wKing := b.getKing(White)
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force pawn into a spot where white king is close to being in check.
	bPawn := b.getPieceByCoordinates(3, 7)
	bPawn.forceMove(&Square{x: 3, y: 3})

	// White's not in check, RIGHT?  Pawns can't move backwards.
	gameState = b.getGameState()
	if gameState == WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardDoesNotRecognizeBlackInCheckByBackwardsPawn(t *testing.T) {
	b := NewBoard()

	// Force the king to the middle of the board.
	bKing := b.getKing(Black)
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force pawn into a spot where king is close to being in check.
	wPawn := b.getPieceByCoordinates(6, 2)
	wPawn.forceMove(&Square{x: 3, y: 5})

	// Black's not in check, RIGHT?  Pawns can't move backwards.
	gameState = b.getGameState()
	if gameState == BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardRecognizesWhiteInCheckByKnight(t *testing.T) {
	b := NewBoard()

	// Force the white king to the middle of the board.
	wKing := b.getKing(White)
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force knight into a spot where white king is in check.
	bKnight := b.getPieceByCoordinates(2, 8)
	bKnight.forceMove(&Square{x: 2, y: 5})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByKnight(t *testing.T) {
	b := NewBoard()

	// Force the king to the middle of the board.
	bKing := b.getKing(Black)
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force knight into a spot where king is in check.
	wKnight := b.getPieceByCoordinates(7, 1)
	wKnight.forceMove(&Square{x: 5, y: 6})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestNewBoardPopulatesCorrectRooks(t *testing.T) {
	b := NewBoard()
	rook_squares := []Square{Square{x: 1, y: 1}, Square{x: 8, y: 1}, Square{x: 1, y: 8}, Square{x: 8, y: 8}}
	for _, square := range rook_squares {
		matching_piece := b.squares[square.x][square.y]
		if !(matching_piece.x() == square.x && matching_piece.y() == square.y) {
			t.Errorf("Piece at index %d, %d reports position as %d, %d", square.x, square.y, matching_piece.x(), matching_piece.y())
		}
		if matching_piece.getShorthand() != "R" {
			t.Errorf("Piece is supposed to be a rook!")
		}
		if square.y == 1 {
			if matching_piece.color != White {
				t.Errorf("Piece is supposed to be white!")
			}
		} else {
			if matching_piece.color != Black {
				t.Errorf("Piece is supposed to be black!")
			}
		}
	}
}

func TestNewBoardPopulatesCorrectKnights(t *testing.T) {
	b := NewBoard()
	knight_squares := []Square{Square{x: 2, y: 1}, Square{x: 7, y: 1}, Square{x: 2, y: 8}, Square{x: 7, y: 8}}
	for _, square := range knight_squares {
		matching_piece := b.squares[square.x][square.y]
		if !(matching_piece.x() == square.x && matching_piece.y() == square.y) {
			t.Errorf("Piece at index %d, %d reports position as %d, %d", square.x, square.y, matching_piece.x(), matching_piece.y())
		}
		if matching_piece.getShorthand() != "N" {
			t.Errorf("Piece is supposed to be a knight!")
		}
		if square.y == 1 {
			if matching_piece.color != White {
				t.Errorf("Piece is supposed to be white!")
			}
		} else {
			if matching_piece.color != Black {
				t.Errorf("Piece is supposed to be black!")
			}
		}
	}
}

func TestNewBoardPopulatesCorrectBishops(t *testing.T) {
	b := NewBoard()
	squares := []Square{Square{x: 3, y: 1}, Square{x: 6, y: 1}, Square{x: 3, y: 8}, Square{x: 6, y: 8}}
	for _, square := range squares {
		matching_piece := b.squares[square.x][square.y]
		if !(matching_piece.x() == square.x && matching_piece.y() == square.y) {
			t.Errorf("Piece at index %d, %d reports position as %d, %d", square.x, square.y, matching_piece.x(), matching_piece.y())
		}
		if matching_piece.getShorthand() != "B" {
			t.Errorf("Piece is supposed to be a bishop!")
		}
		if square.y == 1 {
			if matching_piece.color != White {
				t.Errorf("Piece is supposed to be white!")
			}
		} else {
			if matching_piece.color != Black {
				t.Errorf("Piece is supposed to be black!")
			}
		}
	}
}

func TestNewBoardPopulatesCorrectKingsAndQueens(t *testing.T) {
	b := NewBoard()
	squares := []Square{Square{x: 4, y: 1}, Square{x: 5, y: 1}, Square{x: 4, y: 8}, Square{x: 5, y: 8}}
	for _, square := range squares {
		matching_piece := b.squares[square.x][square.y]
		if !(matching_piece.x() == square.x && matching_piece.y() == square.y) {
			t.Errorf("Piece at index %d, %d reports position as %d, %d", square.x, square.y, matching_piece.x(), matching_piece.y())
		}
		if square.x == 5 {
			if matching_piece.getShorthand() != "K" {
				t.Errorf("Piece is supposed to be a king!")
			}
		} else {
			if matching_piece.getShorthand() != "Q" {
				t.Errorf("Piece is supposed to be a queen!")
			}
		}
		if square.y == 1 {
			if matching_piece.color != White {
				t.Errorf("Piece is supposed to be white!")
			}
		} else {
			if matching_piece.color != Black {
				t.Errorf("Piece is supposed to be black!")
			}
		}
	}
}

func TestNewBoardPopulatesCorrectPawns(t *testing.T) {
	b := NewBoard()
	squares := []Square{Square{x: 1, y: 2}, Square{x: 2, y: 2}, Square{x: 3, y: 2}, Square{x: 4, y: 2}, Square{x: 5, y: 2}, Square{x: 6, y: 2}, Square{x: 7, y: 2}, Square{x: 8, y: 2}, Square{x: 1, y: 7}, Square{x: 2, y: 7}, Square{x: 3, y: 7}, Square{x: 4, y: 7}, Square{x: 5, y: 7}, Square{x: 6, y: 7}, Square{x: 7, y: 7}, Square{x: 8, y: 7}}
	for _, square := range squares {
		matching_piece := b.squares[square.x][square.y]
		if !(matching_piece.x() == square.x && matching_piece.y() == square.y) {
			t.Errorf("Piece at index %d, %d reports position as %d, %d", square.x, square.y, matching_piece.x(), matching_piece.y())
		}
		if matching_piece.getShorthand() != "P" {
			t.Errorf("Piece is supposed to be a pawn!")
		}
		if square.y == 2 {
			if matching_piece.color != White {
				t.Errorf("Piece is supposed to be white!")
			}
		} else {
			if matching_piece.color != Black {
				t.Errorf("Piece is supposed to be black!")
			}
		}
	}
}

func TestNewBoardPopulatesColoredTrackingArrays(t *testing.T) {
	b := NewBoard()
	// Iterate through the arrays which track pieces by color.
	// Is everything there?
	for i, wPiece := range b.colorPieces[White] {
		x := (i % endSquare)
		y := 0
		if i >= endSquare {
			y = 1
		}
		mPiece := b.squares[x+1][y+1]
		if wPiece != mPiece {
			t.Errorf("Piece at index %d is at %d, %d while matching piece is at %d, %d.", i, wPiece.x(), wPiece.y(), mPiece.x(), mPiece.y())
		}
	}
	// Note that the ordering in the tracking arrays matches
	// piece location as you go up the y axis.  That's why black
	// has pawns ordered first in its tracking array.
	for i, bPiece := range b.colorPieces[Black] {
		x := (i % endSquare)
		y := 7
		if i >= endSquare {
			y = 6
		}
		mPiece := b.squares[x+1][y+1]
		if bPiece != mPiece {
			t.Errorf("Piece at index %d is at %d, %d while matching piece is at %d, %d.", i, bPiece.x(), bPiece.y(), mPiece.x(), mPiece.y())
		}
	}
}

func BenchmarkNewBoard(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewBoard()
	}
}

func TestBoardRecognizesSimpleCheckmateOnWhite(t *testing.T) {
	b := NewBoard()

	// Force the white king to the side of the board.
	wKing := b.getKing(White)
	wKing.forceMove(&Square{x: 8, y: 3})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where white king is in check.
	bRook := b.getPieceByCoordinates(8, 8)
	bRook.forceMove(&Square{x: 8, y: 5})

	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}

	// Now king is checkmated.
	bRook2 := b.getPieceByCoordinates(1, 8)
	bRook2.forceMove(&Square{x: 7, y: 6})

	// White's mated, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteCheckmated {
		t.Errorf("State of game should be white is checkmated, but is %d", gameState)
	}
}

func TestBoardRecognizesSimpleCheckmateOnBlack(t *testing.T) {
	b := NewBoard()

	// Force the black king to the side of the board.
	bKing := b.getKing(Black)
	bKing.forceMove(&Square{x: 1, y: 6})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where black king is in check.
	wRook := b.getPieceByCoordinates(1, 1)
	wRook.forceMove(&Square{x: 1, y: 3})

	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}

	// Now king is checkmated.
	wRook2 := b.getPieceByCoordinates(8, 1)
	wRook2.forceMove(&Square{x: 2, y: 3})

	// Black's mated, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackCheckmated {
		t.Errorf("State of game should be black is checkmated, but is %d", gameState)
	}
}

func TestBoardRecognizesBlockedMateOnWhite(t *testing.T) {
	b := NewBoard()

	// Force the white king to the side of the board.
	wKing := b.getKing(White)
	wKing.forceMove(&Square{x: 8, y: 3})

	// Force the bishop into a spot where it can block the check.
	wBishop := b.getPieceByCoordinates(6, 1)
	wBishop.forceMove(&Square{x: 6, y: 3})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where white king is in check.
	bRook := b.getPieceByCoordinates(8, 8)
	bRook.forceMove(&Square{x: 8, y: 5})

	// Now, if the bishop weren't there, the king is checkmated.  But the bishop is there.
	bRook2 := b.getPieceByCoordinates(1, 8)
	bRook2.forceMove(&Square{x: 7, y: 6})

	// White's only in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white is checked, but is %d", gameState)
	}
}

func TestBoardRecognizesBlockedMateOnBlack(t *testing.T) {
	b := NewBoard()

	// Force the black king to the side of the board.
	bKing := b.getKing(Black)
	bKing.forceMove(&Square{x: 1, y: 6})

	// Force the bishop into a spot where it can block the check.
	bBishop := b.getPieceByCoordinates(6, 8)
	bBishop.forceMove(&Square{x: 3, y: 6})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where black king is in check.
	wRook := b.getPieceByCoordinates(1, 1)
	wRook.forceMove(&Square{x: 1, y: 4})

	// Now, if the bishop weren't there, the king is checkmated.  But the bishop is there.
	wRook2 := b.getPieceByCoordinates(8, 8)
	wRook2.forceMove(&Square{x: 2, y: 5})

	// Black's only in check, SI?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black is checked, but is %d", gameState)
	}
}

func TestBoardRejectsAttemptedUserMoveToBadSquare(t *testing.T) {
	b := NewBoard()

	// Bad square, based on vertical index.
	badMove := "Na0"
	success, err := b.attemptUserMove(badMove)
	if success || err == nil {
		t.Errorf("Bad user move allowed to %s.", badMove)
	}

	// Bad square, based on horizontal index.
	badMove = "Km6"
	success, err = b.attemptUserMove(badMove)
	if success || err == nil {
		t.Errorf("Bad user move allowed to %s.", badMove)
	}

	// Good square, bad piece.
	badMove = "Za1"
	success, err = b.attemptUserMove(badMove)
	if success || err == nil {
		t.Errorf("Bad user move allowed to %s.", badMove)
	}

	// Crazy jibber jabber.
	badMove = "barf"
	success, err = b.attemptUserMove(badMove)
	if success || err == nil {
		t.Errorf("Bad user move allowed to %s.", badMove)
	}

	// Crazy jibber jabber, round 2.
	badMove = "33333"
	success, err = b.attemptUserMove(badMove)
	if success || err == nil {
		t.Errorf("Bad user move allowed to %s.", badMove)
	}
	// Empty string, clearly.
	badMove = ""
	success, err = b.attemptUserMove(badMove)
	if success || err == nil {
		t.Errorf("Bad user move allowed to %s.", badMove)
	}

	// TODO: Don't need to specify piece for pawns.
}
