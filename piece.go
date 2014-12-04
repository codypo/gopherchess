package main

type Piece interface {
	move() bool
}

type PieceData struct {
	color    Color
	x        int
	y        int
	captured bool
}

type Pawn struct {
	data PieceData
}

func (p Pawn) move() bool {
	return false
}

type Knight struct {
	data PieceData
}

func (k Knight) move() bool {
	return false
}

type Bishop struct {
	data PieceData
}

func (b Bishop) move() bool {
	return false
}

type Rook struct {
	data PieceData
}

func (r Rook) move() bool {
	return false
}

type Queen struct {
	data PieceData
}

func (q Queen) move() bool {
	return false
}

type King struct {
	data PieceData
}

func (k King) move() bool {
	return false
}
