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
	pawnRow := startSquare + 1
	rookRow := startSquare
	if color == Black {
		pawnRow = endSquare - 1
		rookRow = endSquare
	}

	// Populate pawns.
	for x := startSquare; x <= endSquare; x++ {
		s := Square{x: x, y: pawnRow}
		data := PieceData{color: color, square: s, captured: false}
		pawn := Pawn{data: data}
		p.pieces[pieceIndex] = pawn
		pieceIndex++
	}

	// Populate rooks.
	rook_square1 := Square{x: startSquare, y: rookRow}
	rook_data1 := PieceData{color: color, square: rook_square1, captured: false}
	rook1 := Rook{data: rook_data1}
	p.pieces[pieceIndex] = rook1
	pieceIndex++

	// TODO: This pieceIndex part is silly.
	rook_square2 := Square{x: endSquare, y: rookRow}
	rook_data2 := PieceData{color: color, square: rook_square2, captured: false}
	rook2 := Rook{data: rook_data2}
	p.pieces[pieceIndex] = rook2
	pieceIndex++

	return p
}
