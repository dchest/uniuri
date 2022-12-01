// Written in 2011-2014 by Dmitry Chestnykh
//
// The author(s) have dedicated all copyright and related and
// neighboring rights to this software to the public domain
// worldwide. Distributed without any warranty.
// http://creativecommons.org/publicdomain/zero/1.0/

// Package uniuri generates random strings good for use in URIs to identify
// unique objects.
//
// Example usage:
//
//	s := uniuri.New() // s is now "apHCJBl7L1OmC57n"
//
// A standard string created by New() is 16 bytes in length and consists of
// Latin upper and lowercase letters, and numbers (from the set of 62 allowed
// characters), which means that it has ~95 bits of entropy. To get more
// entropy, you can use NewLen(UUIDLen), which returns 20-byte string, giving
// ~119 bits of entropy, or any other desired length.
//
// Functions read from crypto/rand random source, and panic if they fail to
// read from it.
package uniuri

import (
	"crypto/rand"
	"math"
)

const (
	// StdLen is a standard length of uniuri string to achive ~95 bits of entropy.
	StdLen = 16
	// UUIDLen is a length of uniuri string to achive ~119 bits of entropy, closest
	// to what can be losslessly converted to UUIDv4 (122 bits).
	UUIDLen = 20
)

// StdChars is a set of standard characters allowed in uniuri string.
// Deprecated: use functional options instead of slice modification
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// New returns a new random string of the standard length, consisting of
// standard characters if no custom options given.
func New(opts ...Opt) string {
	return string(NewBytes(opts...))
}

// NewBytes returns a new random byte slice of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
func NewBytes(opts ...Opt) []byte {
	o := options{
		length: StdLen,
		chars:  StdChars,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return generate(o.length, o.chars)
}

// NewLen returns a new random string of the provided length, consisting of
// standard characters.
// Deprecated: Use New with proper options instead.
func NewLen(length int) string {
	return New(Length(length))
}

// NewLenChars returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
// Deprecated: Use New with proper options instead.
func NewLenChars(length int, chars []byte) string {
	return string(New(Length(length), Chars(chars)))
}

// NewLenCharsBytes returns a new random byte slice of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
// Deprecated: Use NewBytes with proper options instead.
func NewLenCharsBytes(length int, chars []byte) []byte {
	return NewBytes(Length(length), Chars(chars))
}

// maxBufLen is the maximum length of a temporary buffer for random bytes.
const maxBufLen = 2048

// minRegenBufLen is the minimum length of temporary buffer for random bytes
// to fill after the first rand.Read request didn't produce the full result.
// If the initial buffer is smaller, this value is ignored.
// Rationale: for performance, assume it's pointless to request fewer bytes from rand.Read.
const minRegenBufLen = 16

// estimatedBufLen returns the estimated number of random bytes to request
// given that byte values greater than maxByte will be rejected.
func estimatedBufLen(need, maxByte int) int {
	return int(math.Ceil(float64(need) * (255 / float64(maxByte))))
}

// generate creates a new random byte slice of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
func generate(length int, chars []byte) []byte {
	if length == 0 {
		return nil
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("uniuri: wrong charset length for NewLenChars")
	}
	maxrb := 255 - (256 % clen)
	buflen := estimatedBufLen(length, maxrb)
	if buflen < length {
		buflen = length
	}
	if buflen > maxBufLen {
		buflen = maxBufLen
	}
	buf := make([]byte, buflen) // storage for random bytes
	out := make([]byte, length) // storage for result
	i := 0
	for {
		if _, err := rand.Read(buf[:buflen]); err != nil {
			panic("uniuri: error reading random bytes: " + err.Error())
		}
		for _, rb := range buf[:buflen] {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			out[i] = chars[c%clen]
			i++
			if i == length {
				return out
			}
		}
		// Adjust new requested length, but no smaller than minRegenBufLen.
		buflen = estimatedBufLen(length-i, maxrb)
		if buflen < minRegenBufLen && minRegenBufLen < cap(buf) {
			buflen = minRegenBufLen
		}
		if buflen > maxBufLen {
			buflen = maxBufLen
		}
	}
}

// options is a set of variables to be used to generate random values
type options struct {
	length int
	chars  []byte
}

// Opt is a function to set specific value to options
type Opt func(opts *options)

// Length sets length to options
func Length(length int) Opt {
	return func(opts *options) {
		opts.length = length
	}
}

// Chars sets chars slice to options
func Chars(chars []byte) Opt {
	return func(opts *options) {
		opts.chars = chars
	}
}
