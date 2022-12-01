// Written in 2011-2014 by Dmitry Chestnykh
//
// The author(s) have dedicated all copyright and related and
// neighboring rights to this software to the public domain
// worldwide. Distributed without any warranty.
// http://creativecommons.org/publicdomain/zero/1.0/

package uniuri

import (
	"testing"
)

var (
	testCharSet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

func Test_New(t *testing.T) {
	testCases := []struct {
		name           string
		options        []Opt
		expectedLength int
		expectedChars  []byte
	}{
		{
			name:           "default",
			expectedLength: 16,
			expectedChars:  testCharSet,
		},
		{
			name:           "with_length",
			options:        []Opt{Length(24)},
			expectedLength: 24,
			expectedChars:  testCharSet,
		},
		{
			name:           "with_chars",
			options:        []Opt{Chars(threeChars)},
			expectedLength: 16,
			expectedChars:  threeChars,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := New(tc.options...)

			if len(u) != tc.expectedLength {
				t.Errorf("wrong length: expected %d, got %d", tc.expectedLength, len(u))
			}

			allowedChars := make(map[rune]struct{})
			for _, c := range tc.expectedChars {
				allowedChars[rune(c)] = struct{}{}
			}
			for _, r := range []rune(u) {
				if _, ok := allowedChars[r]; !ok {
					t.Errorf("character not allowed: %s", string(r))
				}
			}
		})
	}
}

func Test_NewBytes(t *testing.T) {
	testCases := []struct {
		name           string
		options        []Opt
		expectedLength int
		expectedChars  []byte
	}{
		{
			name:           "default",
			expectedLength: 16,
			expectedChars:  testCharSet,
		},
		{
			name:           "with_length",
			options:        []Opt{Length(24)},
			expectedLength: 24,
			expectedChars:  testCharSet,
		},
		{
			name:           "with_chars",
			options:        []Opt{Chars(threeChars)},
			expectedLength: 16,
			expectedChars:  threeChars,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := NewBytes(tc.options...)

			if len(u) != tc.expectedLength {
				t.Errorf("wrong length: expected %d, got %d", tc.expectedLength, len(u))
			}

			allowedChars := make(map[rune]struct{})
			for _, c := range tc.expectedChars {
				allowedChars[rune(c)] = struct{}{}
			}
			for _, b := range u {
				if _, ok := allowedChars[rune(b)]; !ok {
					t.Errorf("character not allowed: %s", string(b))
				}
			}
		})
	}
}

func Test_unique(t *testing.T) {
	uris := make(map[string]struct{})
	for i := 0; i < 10000; i++ {
		u := New()
		if _, ok := uris[u]; ok {
			t.Fatalf("non-unique uniuri: %s", u)
		}
		uris[u] = struct{}{}
	}
}

func Test_bias(t *testing.T) {
	chars := []byte("abcdefghijklmnopqrstuvwxyz")
	slen := 100000
	s := New(Length(slen), Chars(chars))
	counts := make(map[rune]int)
	for _, b := range s {
		counts[b]++
	}
	avg := float64(slen) / float64(len(chars))
	for k, n := range counts {
		diff := float64(n) / avg
		if diff < 0.95 || diff > 1.05 {
			t.Errorf("Possible bias on '%c': expected average %f, got %d", k, avg, n)
		}
	}
}
