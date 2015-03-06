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

func TestBoardRecognizesWhiteInCheck(t *testing.T) {
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

func TestBoardRecognizesBlackInCheck(t *testing.T) {
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

	// White's in check, RIGHT?
	gameState = b.getGameState()
	if gameState != BlackInCheck {
		t.Errorf("State of game should be black in check, but is %d", gameState)
	}
}
