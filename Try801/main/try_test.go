package main

import "testing"

func TestIsEven(t *testing.T) {
	// テストケースを作成
	cases := map[string]struct {
		in   int
		want bool
	}{
		"+odd":  {in: 5, want: false},
		"+even": {in: 6, want: true},
		"-odd":  {in: -5, want: false},
		"-even": {in: -6, want: true},
		"+zero": {in: 0, want: true},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := IsEven(tt.in); tt.want != got {
				t.Errorf("want IsOdd(%d) = %v, got%v", tt.in, tt.want, got)
			}
		})
	}
}

func IsEven(n int) bool {
	return n % 2 != 0
}