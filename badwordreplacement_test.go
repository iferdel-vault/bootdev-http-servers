package main

import "testing"

func TestBadWordReplacement(t *testing.T) {
	tests := map[string]struct {
		WordsIn  string
		WordsOut string
	}{
		"single word": {
			WordsIn:  "kerfuffle",
			WordsOut: "****",
		},
		"uppercase": {
			WordsIn:  "Kerfuffle",
			WordsOut: "****",
		},
		"punctuation": {
			WordsIn:  "¿Kerfuffle?",
			WordsOut: "¿Kerfuffle?",
		},
		"two bad words in sentence": {
			WordsIn:  "fornax along with kerfuffle",
			WordsOut: "**** along with ****",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := BadWordReplacement(tc.WordsIn)
			if got != tc.WordsOut {
				t.Errorf("BadWordReplacement(%q) --> expected %q, got %q", tc.WordsIn, tc.WordsOut, got)
			}
		})
	}
}
