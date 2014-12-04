package main

type Board struct {
	players [numPlayers]Player
	squares [numSquaresWide][numSquaresTall]Square
}
