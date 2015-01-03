package main

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

func (b Board) evaluateSquare(c Color, s *Square) int {
	if !s.isValid() {
		return squareInvalid
	}
	status := squareVacant
	var me Player
	var opponent Player

	if c == White {
		me = b.getPlayer(White)
		opponent = b.getPlayer(Black)
	} else {
		me = b.getPlayer(Black)
		opponent = b.getPlayer(White)
	}

	// Does either side have a piece on this square?
	myPiece := me.getPieceByCoordinate(s.x, s.y)
	if myPiece != nil {
		status = squareOccupiedByMe
	}

	oppoPiece := opponent.getPieceByCoordinate(s.x, s.y)
	if oppoPiece != nil {
		status = squareOccupiedByOpponent
	}

	// Eventually, we'll want more stuff for check and such.
	return status
}
