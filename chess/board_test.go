package chess

import "testing"

func TestSquare(t *testing.T) {
	tests := []struct {
		file int
		rank int
		want int
	}{
		{file: 0, rank: 0, want: 0},
		{file: 4, rank: 0, want: 4},
		{file: 0, rank: 1, want: 8},
		{file: 7, rank: 7, want: 63},
	}

	for _, tt := range tests {
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
	}
}
