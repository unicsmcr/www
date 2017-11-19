package rand

import (
	"math/rand"
	"time"
)

// A source that is seeded with time.Now().UnixNano() for different outputs at
// different runs.
var src = rand.NewSource(time.Now().UnixNano())

const (
	lettDig      = "0123456789abcdefghijklmnopqrstuvwqxyzABCDEFGHIJKLMNOPQRSTUVQWXYZ"
	letterIDBits = 6                 // 6 bits to represent a letter index
	letterIDMask = 1<<6 - 1          // all the bits set to 1
	letterMax    = 63 / letterIDBits // number of letters that fit in 63 bits
)

// Source is a pseudo-random source that is seeded with the specified value.
type Source interface {
	Int63() int64
	Seed(seed int64)
}

// Src gets the current source.
func Src() Source {
	return src
}

// String generates a random string of length n.
func String(n int) string {
	a := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterMax
		}
		if idx := int(cache & letterIDMask); idx < len(lettDig) {
			a[i] = lettDig[idx]
			i--
		}
		cache >>= letterIDBits
		remain--
	}
	return string(a)
}
