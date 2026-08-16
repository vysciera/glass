package chess

import "testing"

func TestSquare(t *testing.T) {
	tests := []struct {
		file int
		rank int
		want int
	}{
		{0, 0, 0},
		{4, 0, 4},
		{0, 1, 8},
		{7, 7, 63},
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
