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
