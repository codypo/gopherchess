package main

type Bishop struct {
	data PieceData
}

func (b Bishop) move() bool {
	return false
}
