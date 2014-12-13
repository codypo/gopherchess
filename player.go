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

	// Initialize a square and piece data instance we can reuse.
	s := Square{x: startSquare, y: startSquare}
	pd := PieceData{color: color, square: s, captured: false}
	pawn := Pawn{data: pd}

	// Populate pawns.
	for x := startSquare; x <= endSquare; x++ {
		s = Square{x: x, y: pawnRow}
		pd = PieceData{color: color, square: s, captured: false}
		pawn = Pawn{data: pd}
		p.pieces[pieceIndex] = pawn
		pieceIndex++
	}

	// Populate rooks.
	s = Square{x: startSquare, y: rookRow}
	pd = PieceData{color: color, square: s, captured: false}
	rook := Rook{data: pd}
	p.pieces[pieceIndex] = rook
	pieceIndex++

	// TODO: This pieceIndex part is silly.
	s = Square{x: endSquare, y: rookRow}
	pd = PieceData{color: color, square: s, captured: false}
	rook = Rook{data: pd}
	p.pieces[pieceIndex] = rook
	pieceIndex++

	// Populate knights.
	s = Square{x: startSquare + 1, y: rookRow}
	pd = PieceData{color: color, square: s, captured: false}
	knight := Knight{data: pd}
	p.pieces[pieceIndex] = knight
	pieceIndex++

	s = Square{x: endSquare - 1, y: rookRow}
	pd = PieceData{color: color, square: s, captured: false}
	knight = Knight{data: pd}
	p.pieces[pieceIndex] = knight
	pieceIndex++

	// Populate bishops.
	s = Square{x: startSquare + 1, y: rookRow}
	pd = PieceData{color: color, square: s, captured: false}
	bishop := Bishop{data: pd}
	p.pieces[pieceIndex] = bishop
	pieceIndex++

	s = Square{x: endSquare - 1, y: rookRow}
	pd = PieceData{color: color, square: s, captured: false}
	bishop = Bishop{data: pd}
	p.pieces[pieceIndex] = bishop
	pieceIndex++

	return p
}
