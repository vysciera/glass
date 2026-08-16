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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.piece.Symbol()

			if got != tt.want {
				t.Errorf("Symbol() = &q; want %q", got, tt.want)
			}
		})
	}
}
