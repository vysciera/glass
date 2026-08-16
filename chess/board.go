package chess

type Board [64]Piece

func Square(file, rank int) int {
	return rank*8 + file
}

func (b *Board) Set(file, rank int, piece Piece) {
	b[Square(file, rank)] = piece
}
