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

	// Populate knights.
	knight_square1 := Square{x: startSquare + 1, y: rookRow}
	knight_data1 := PieceData{color: color, square: knight_square1, captured: false}
	knight1 := Knight{data: knight_data1}
	p.pieces[pieceIndex] = knight1
	pieceIndex++

	knight_square2 := Square{x: endSquare - 1, y: rookRow}
	knight_data2 := PieceData{color: color, square: knight_square2, captured: false}
	knight2 := Knight{data: knight_data2}
	p.pieces[pieceIndex] = knight2
	pieceIndex++

	// Populate bishops.
	bishop_square1 := Square{x: startSquare + 1, y: rookRow}
	bishop_data1 := PieceData{color: color, square: bishop_square1, captured: false}
	bishop1 := Bishop{data: bishop_data1}
	p.pieces[pieceIndex] = bishop1
	pieceIndex++

	bishop_square2 := Square{x: endSquare - 1, y: rookRow}
	bishop_data2 := PieceData{color: color, square: bishop_square2, captured: false}
	bishop2 := Bishop{data: bishop_data2}
	p.pieces[pieceIndex] = bishop2
	pieceIndex++

	return p
}
