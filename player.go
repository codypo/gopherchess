package main

type Player struct {
	name   string
	color  Color
	pieces [numPiecesPerPlayer]Piece
}

func NewPlayer(color Color) *Player {
	p := new(Player)
	p.color = color

	// Piece initialization goes here.
	pieceIndex := 0

	// p p p p p p p p.
	// r k b q k b k r.
	pawnRow := 2
	if color == Black {
		pawnRow = 7
	}

	for x := startSquare; x <= endSquare; x++ {
		s := Square{x: x, y: pawnRow}
		data := PieceData{color: color, square: s, captured: false}
		pawn := Pawn{data: data}
		p.pieces[pieceIndex] = pawn
		pieceIndex++
	}

	return p
}
