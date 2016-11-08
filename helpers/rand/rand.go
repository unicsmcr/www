package rand

import (
	"math/rand"
	"time"
)

// new source for rand func, seeding it with time.Now().UnixNano()
// for different outputs at different runs
var src = rand.NewSource(time.Now().UnixNano())

const (
	lettDig      = "0123456789abcdefghijklmnopqrstuvwqxyzABCDEFGHIJKLMNOPQRSTUVQWXYZ"
	letterIdBits = 6                 // 6 bits to represent a letter index
	letterIdMask = 1<<6 - 1          // all the bits set to 1
	letterMax    = 63 / letterIdBits // number of letter indices fitting in 63 bits
)

// must define type Source
type Source interface {
	Int63() int64
	Seed(seed int64)
}

// getter for our new source
func Src() Source {
	return src
}

// rand function, using byte optimization, using almost all the
// 64 bits of the src.Int63() (by shifting them), in order to
// minimize the number of calls
func RandString(n int) string {
	a := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterMax
		}
		if idx := int(cache & letterIdMask); idx < len(lettDig) {
			a[i] = lettDig[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return string(a)
}
