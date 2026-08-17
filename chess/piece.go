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

func (p Piece) Symbol() string {
	switch p.Type {
	case NoPiece:
		return "."

	case King:
		if p.Color == White {
			return "K"
		}
		return "k"

	case Queen:
		if p.Color == White {
			return "Q"
		}
		return "q"

	case Bishop:
		if p.Color == White {
			return "B"
		}
		return "b"

	case Knight:
		if p.Color == White {
			return "N"
		}
		return "n"

	case Rook:
		if p.Color == White {
			return "R"
		}
		return "r"

	case Pawn:
		if p.Color == White {
			return "P"
		}
		return "p"
	}

	return "?"
}
