// Package chess manages board position & engine logic
package chess

import "strings"

type Board [64]Piece

func Square(file, rank int) int {
	return rank*8 + file
}

func (b *Board) Set(file, rank int, piece Piece) {
	b[Square(file, rank)] = piece
}

func NewStartingBoard() Board {
	var board Board

	for file := 0; file < 8; file++ {
		board.Set(file, 1, Piece{
			Type:  Pawn,
			Color: White,
		})

		board.Set(file, 6, Piece{
			Type:  Pawn,
			Color: Black,
		})
	}

	backRank := [8]PieceType{
		Rook,
		Knight,
		Bishop,
		Queen,
		King,
		Bishop,
		Knight,
		Rook,
	}

	for file, pieceType := range backRank {
		board.Set(file, 0, Piece{
			Type:  pieceType,
			Color: White,
		})

		board.Set(file, 7, Piece{
			Type:  pieceType,
			Color: Black,
		})
	}

	return board
}

func (b Board) String() string {
	var out strings.Builder

	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			symbol := b[Square(file, rank)].Symbol()

			if file != 7 {
				symbol += " "
			}

			out.WriteString(symbol)
		}

		if rank != 0 {
			out.WriteString("\n")
		}
	}

	return out.String()
}
