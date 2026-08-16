// Package chess manages board position & engine logic
package chess

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
