package main

import (
	"testing"
)

func TestBoardNew(t *testing.T) {
	b := NewBoard()
	totalSquares := 0
	for i, _ := range b.squares {
		for j, _ := range b.squares[i] {
			totalSquares++

			s := b.squares[i][j]
			if !s.isValid() {
				t.Errorf("Board contains invalid square at %d, %d", s.x, s.y)
			}
		}
	}
	if totalSquares != 64 {
		t.Errorf("Board contains %d squares.", totalSquares)
	}

	if len(b.players) != 2 {
		t.Errorf("Board contains %d players.", len(b.players))
	}
	if b.players[0].color != White {
		t.Errorf("Player 0 is not White.")
	}
	if b.players[1].color != Black {
		t.Errorf("Player 1 is not Black.")
	}
}

func TestBoardRecognizesWhiteInCheckByRook(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the white king to the middle of the board.
	wKing, _ := white.getKing()
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where white king is in check.
	bRook, _ := black.getPieceByCoordinate(8, 8)
	bRook.forceMove(&Square{x: 8, y: 4})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByRook(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the king to the middle of the board.
	bKing, _ := black.getKing()
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where king is in check.
	wRook, _ := white.getPieceByCoordinate(1, 1)
	wRook.forceMove(&Square{x: 1, y: 4})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardRecognizesWhiteInCheckByQueen(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the white king to the middle of the board.
	wKing, _ := white.getKing()
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force queen into a spot where white king is in check.
	bQueen, _ := black.getPieceByCoordinate(4, 8)
	bQueen.forceMove(&Square{x: 6, y: 6})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByQueen(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the king to the middle of the board.
	bKing, _ := black.getKing()
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force rook into a spot where king is in check.
	wQueen, _ := white.getPieceByCoordinate(4, 1)
	wQueen.forceMove(&Square{x: 3, y: 3})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardRecognizesWhiteInCheckByBishop(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the white king to the middle of the board.
	wKing, _ := white.getKing()
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force bishop into a spot where white king is in check.
	bBishop, _ := black.getPieceByCoordinate(3, 8)
	bBishop.forceMove(&Square{x: 6, y: 6})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByBishop(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the king to the middle of the board.
	bKing, _ := black.getKing()
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force bishop into a spot where king is in check.
	wBishop, _ := white.getPieceByCoordinate(6, 1)
	wBishop.forceMove(&Square{x: 5, y: 5})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardRecognizesWhiteInCheckByPawn(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the white king to the middle of the board.
	wKing, _ := white.getKing()
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force pawn into a spot where white king is in check.
	bPawn, _ := black.getPieceByCoordinate(3, 7)
	bPawn.forceMove(&Square{x: 3, y: 5})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByPawn(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the king to the middle of the board.
	bKing, _ := black.getKing()
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force pawn into a spot where king is in check.
	wPawn, _ := white.getPieceByCoordinate(6, 2)
	wPawn.forceMove(&Square{x: 3, y: 3})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardDoesNotRecognizeWhiteInCheckByBackwardsPawn(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the white king to the middle of the board.
	wKing, _ := white.getKing()
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force pawn into a spot where white king is close to being in check.
	bPawn, _ := black.getPieceByCoordinate(3, 7)
	bPawn.forceMove(&Square{x: 3, y: 3})

	// White's not in check, RIGHT?  Pawns can't move backwards.
	gameState = b.getGameState()
	if gameState == WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardDoesNotRecognizeBlackInCheckByBackwardsPawn(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the king to the middle of the board.
	bKing, _ := black.getKing()
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force pawn into a spot where king is close to being in check.
	wPawn, _ := white.getPieceByCoordinate(6, 2)
	wPawn.forceMove(&Square{x: 3, y: 5})

	// Black's not in check, RIGHT?  Pawns can't move backwards.
	gameState = b.getGameState()
	if gameState == BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}

func TestBoardRecognizesWhiteInCheckByKnight(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the white king to the middle of the board.
	wKing, _ := white.getKing()
	wKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force knight into a spot where white king is in check.
	bKnight, _ := black.getPieceByCoordinate(2, 8)
	bKnight.forceMove(&Square{x: 2, y: 5})

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != WhiteInCheck {
		t.Errorf("State of game should be white in check, but is %d", gameState)
	}
}

func TestBoardRecognizesBlackInCheckByKnight(t *testing.T) {
	b := NewBoard()
	white := b.getPlayer(White)
	black := b.getPlayer(Black)

	// Force the king to the middle of the board.
	bKing, _ := black.getKing()
	bKing.forceMove(&Square{x: 4, y: 4})

	// All good.
	gameState := b.getGameState()
	if gameState != GameOn {
		t.Errorf("State of game should be default, but is %d", gameState)
	}

	// Force knight into a spot where king is in check.
	wKnight, _ := white.getPieceByCoordinate(7, 1)
	wKnight.forceMove(&Square{x: 5, y: 6})

	// Black's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}
