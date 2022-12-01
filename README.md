Package uniuri
=====================

```go
import "github.com/dchest/uniuri"
```

Package uniuri generates random strings good for use in URIs to identify
unique objects.

Example usage:

```go
s := uniuri.New() // s is now "apHCJBl7L1OmC57n"
```

A standard string created by New() is 16 bytes in length and consists of
Latin upper and lowercase letters, and numbers (from the set of 62 allowed
characters), which means that it has ~95 bits of entropy. To get more
entropy, you can use NewLen(UUIDLen), which returns 20-byte string, giving
~119 bits of entropy, or any other desired length.

Functions read from crypto/rand random source, and panic if they fail to
read from it.


Constants
---------

```go
const (
// StdLen is a standard length of uniuri string to achive ~95 bits of entropy.
StdLen = 16
// UUIDLen is a length of uniuri string to achive ~119 bits of entropy, closest
// to what can be losslessly converted to UUIDv4 (122 bits).
UUIDLen = 20
)

```

Variables
---------

```go
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
```

StdChars is a set of standard characters allowed in uniuri string. It will be used by default by `New` and `NewBytes`
functions unless option `Chars` given.


Functions
---------

### func New

```go
func New(opts ...Opt) string
```

New returns a new random string. It accepts zero or more functional options to change output.

### func NewLen

```go
func NewBytes(opts ...Opt) string
```

New returns a new slice of random bytes. It accepts zero or more functional options to change output.

Options
---------

### func Length

```go
func Length(length int) Opt
```

Sets length of random string or slice of bytes to be generated.

### func Chars

```go
func Chars(chars []byte) Opt
```

Sets allowed characters to be used upon random string or slice of bytes generation.


Public domain dedication
------------------------

Written in 2011-2014 by Dmitry Chestnykh

The author(s) have dedicated all copyright and related and
neighboring rights to this software to the public domain
worldwide. Distributed without any warranty.
http://creativecommons.org/publicdomain/zero/1.0/

