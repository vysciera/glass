package chess

import "testing"

func TestPieceSymbol(t *testing.T) {
	tests := []struct {
		name  string
		piece Piece
		want  string
	}{
		{
			name: "white king",
			piece: Piece{
				Type:  King,
				Color: White,
			},
			want: "K",
		},
		{
			name: "black king",
			piece: Piece{
				Type:  King,
				Color: Black,
			},
			want: "k",
		},
		{
			name: "empty square",
			piece: Piece{
				Type: NoPiece,
			},
			want: ".",
		},
		{
			name: "white queen",
			piece: Piece{
				Type:  Queen,
				Color: White,
			},
			want: "Q",
		},
		{
			name: "black queen",
			piece: Piece{
				Type:  Queen,
				Color: Black,
			},
			want: "q",
		},
		{
			name: "white knight",
			piece: Piece{
				Type:  Knight,
				Color: White,
			},
			want: "N",
		},
		{
			name: "black knight",
			piece: Piece{
				Type:  Knight,
				Color: Black,
			},
			want: "n",
		},
		{
			name: "white pawn",
			piece: Piece{
				Type:  Pawn,
				Color: White,
			},
			want: "P",
		},
		{
			name: "black pawn",
			piece: Piece{
				Type:  Pawn,
				Color: Black,
			},
			want: "p",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.piece.Symbol()

			if got != tt.want {
				t.Errorf("Symbol() = %q; want %q", got, tt.name)
			}
		})
	}
}
