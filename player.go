package main

import "fmt"

type Player struct {
	name   string
	color  Color
	pieces [numPiecesPerPlayer]Piece
}

func NewPlayer(color Color, board *Board) *Player {
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
	s := &Square{x: startSquare, y: startSquare}
	pd := NewPieceData(color, s, board)
	pawn := Pawn{data: pd}

	// Populate pawns.
	for x := startSquare; x <= endSquare; x++ {
		s = &Square{x: x, y: pawnRow}
		pd = NewPieceData(color, s, board)
		pawn = Pawn{data: pd}
		p.pieces[pieceIndex] = pawn
		pieceIndex++
	}

	// Populate rooks.
	s = &Square{x: startSquare, y: rookRow}
	pd = NewPieceData(color, s, board)
	rook := Rook{data: pd}
	p.pieces[pieceIndex] = rook
	pieceIndex++

	// TODO: This pieceIndex part is silly.
	s = &Square{x: endSquare, y: rookRow}
	pd = NewPieceData(color, s, board)
	rook = Rook{data: pd}
	p.pieces[pieceIndex] = rook
	pieceIndex++

	// Populate knights.
	s = &Square{x: startSquare + 1, y: rookRow}
	pd = NewPieceData(color, s, board)
	knight := Knight{data: pd}
	p.pieces[pieceIndex] = knight
	pieceIndex++

	s = &Square{x: endSquare - 1, y: rookRow}
	pd = NewPieceData(color, s, board)
	knight = Knight{data: pd}
	p.pieces[pieceIndex] = knight
	pieceIndex++

	// Populate bishops.
	s = &Square{x: startSquare + 2, y: rookRow}
	pd = NewPieceData(color, s, board)
	bishop := Bishop{data: pd}
	p.pieces[pieceIndex] = bishop
	pieceIndex++

	s = &Square{x: endSquare - 2, y: rookRow}
	pd = NewPieceData(color, s, board)
	bishop = Bishop{data: pd}
	p.pieces[pieceIndex] = bishop
	pieceIndex++

	// Populate the queen.
	s = &Square{x: startSquare + 3, y: rookRow}
	pd = NewPieceData(color, s, board)
	queen := Queen{data: pd}
	p.pieces[pieceIndex] = queen
	pieceIndex++

	// Populate the king.
	s = &Square{x: startSquare + 4, y: rookRow}
	pd = NewPieceData(color, s, board)
	king := King{data: pd}
	p.pieces[pieceIndex] = king
	pieceIndex++

	return p
}

func (player Player) getPieceByCoordinate(x int, y int) (Piece, error) {
	for _, piece := range player.pieces {
		if piece.pieceData().matchesCoordinates(x, y) {
			return piece, nil
		}
		//fmt.Printf("   Piece coords %d, %d does not match %d, %d\n", piece.pieceData().getSquare().x, piece.pieceData().getSquare().y, x, y)
	}

	return nil, fmt.Errorf("No piece found at %d, %d", x, y)
}

func (p Player) printPieces() {
	for _, piece := range p.pieces {
		fmt.Printf("Piece coords %d, %d\n", piece.pieceData().getSquare().x, piece.pieceData().getSquare().y)
	}
}
