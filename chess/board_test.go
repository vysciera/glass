package chess

import "testing"

func TestSquare(t *testing.T) {
	tests := []struct {
		name string
		file int
		rank int
		want int
	}{
		{
			name: "a1 is square 0",
			file: 0,
			rank: 0,
			want: 0,
		},
		{
			name: "e1 is square 4",
			file: 4,
			rank: 0,
			want: 4,
		},
		{
			name: "a2 is square 8",
			file: 0,
			rank: 1,
			want: 8,
		},
		{
			name: "h8 is square 63",
			file: 7,
			rank: 7,
			want: 63,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Square(tt.file, tt.rank)

			if got != tt.want {
				t.Errorf(
					"Square(%d, %d) = %d; want %d",
					tt.file,
					tt.rank,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestBoardSet(t *testing.T) {
	var board Board

	t.Run("set piece at e1", func(t *testing.T) {
		want := Piece{
			Type:  King,
			Color: White,
		}

		board.Set(4, 0, want)
		got := board[Square(4, 0)]

		if got != want {
			t.Errorf("piece at e1 = %+v; want %+v", got, want)
		}
	})
}

func TestNewStartingBoard(t *testing.T) {
	board := NewStartingBoard()

	t.Run("white king starts on e1", func(t *testing.T) {
		want := Piece{
			Type:  King,
			Color: White,
		}

		got := board[Square(4, 0)]

		if got != want {
			t.Errorf("piece at e1 = %+v; want %+v", got, want)
		}
	})

	t.Run("black king starts on e8", func(t *testing.T) {
		want := Piece{
			Type:  King,
			Color: Black,
		}

		got := board[Square(4, 7)]

		if got != want {
			t.Errorf("piece at e8 = %+v; want %+v", got, want)
		}
	})

	t.Run("e4 starts empty", func(t *testing.T) {
		got := board[Square(4, 3)]

		if got.Type != NoPiece {
			t.Errorf("e4 piece type = %+v; want %+v", got.Type, NoPiece)
		}
	})
}
