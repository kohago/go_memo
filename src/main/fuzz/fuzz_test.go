package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := reverse(tc.in)
		if rev != tc.want {
			t.Errorf("want:%v,is:%v", tc.want, rev)
		}
	}
}

//go test -fuzz=Fuzz
//
//A failure occurred while fuzzing, and the input that caused the problem is written to a seed corpus file
//that will be run the next time go test is called,
//even without the -fuzz flag. T
//o view the input that caused the failure, open the corpus file written to the testdata/fuzz/FuzzReverse directory in a text editor.Your seed corpus file may contain a different string, but the format will be the same.
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := reverse(orig)
		doubleRev := reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
