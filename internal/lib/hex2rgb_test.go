package lib

import "testing"

func TestHexToRGB(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{
			// "white"
			in:   "#d4d4d4",
			want: []int{212, 212, 212},
		},
		{
			// "black"
			in:   "#303030",
			want: []int{48, 48, 48},
		},
		{
			// green
			in:   "#608b4e",
			want: []int{96, 139, 78},
		},
		{
			// light blue
			in:   "#9cdcfe",
			want: []int{156, 220, 254},
		},
		{
			// dark blue
			in:   "#569cd6",
			want: []int{86, 156, 214},
		},
	}

	for _, c := range cases {
		r, g, b, _ := HexToRGB(c.in)
		if r != c.want[0] {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, r, c.want[0])
		}
		if g != c.want[1] {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, g, c.want[1])
		}
		if b != c.want[2] {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, b, c.want[2])
		}
	}
}
