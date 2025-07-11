package main

import "testing"

func TestBadWordReplacement(t *testing.T) {
	tests := map[string]struct {
		WordsIn  string
		WordsOut string
	}{
		"right": {
			WordsIn:  "kerfuffle",
			WordsOut: "****",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := BadWordReplacement(tc.WordsIn)
			if got != tc.WordsOut {
				t.Errorf("BadWordReplacement(%v) --> expected %v, got %v", tc.WordsIn, tc.WordsOut, got)
			}
		})
	}
}
