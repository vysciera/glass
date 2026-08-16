package chess

type Color uint8

const ( // ??? Read-up on this later
	White Color = iota
	Black
)

type PieceType uint8

const (
	NoPiece PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Type  PieceType
	Color Color
}
