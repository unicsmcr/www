package validator

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

type testPair struct {
	value  string
	result bool
}

var testsEmail = []testPair{
	{randString(91) + "@" + randString(4) + "." + randString(3), true},
	{randString(50) + "@gmail.com", true},
	{randString(10) + "@" + randString(5) + "." + randString(3), true},
	{randString(20), false},
	{randString(99), false},
	{randString(50) + "@" + randString(10), false},
	{randString(10) + "." + randString(4), false},
	{randString(100), false},
}

var testsMessage = []testPair{
	{randString(3999), true},
	{randString(int(src.Int63() % 400)), true},
	{strings.Repeat(" ", 50) + randString(1), true},
	{strings.Repeat(" ", 50) + randString(3950), true},
	{randString(4001), false},
	{strings.Repeat(" ", 100), false},
}

var testsName = []testPair{
	{randString(29), true},
	{randString(1), true},
	{randString(1) + strings.Repeat(" ", 29), true},
	{randString(1) + strings.Repeat(" ", 30), false},
	{randString(31), false},
	{strings.Repeat(" ", int(src.Int63()%31)), false},
}

func TestIsValidEmail(t *testing.T) {
	for _, pair := range testsEmail {
		res := IsValidEmail(pair.value)
		if res != pair.result {
			t.Error(
				"For", pair.value,
				"expected result", pair.result,
				"got", res,
			)
		}
	}
}

func TestIsValidMessage(t *testing.T) {
	for _, pair := range testsMessage {
		res := IsValidMessage(pair.value)
		if res != pair.result {
			t.Error(
				"For", pair.value,
				"expected result", pair.result,
				"got", res,
			)
		}
	}
}

func TestIsValidName(t *testing.T) {
	for _, pair := range testsName {
		res := IsValidName(pair.value)
		if res != pair.result {
			t.Error(
				"For", pair.value,
				"expected result", pair.result,
				"got", res,
			)
		}
	}
}

// new source for rand func, seeding it with time.Now().UnixNano()
// for different outputs at differents runs
var src = rand.NewSource(time.Now().UnixNano())

const (
	lettDig      = "0123456789abcdefghijklmnopqrstuvwqxyzABCDEFGHIJKLMNOPQRSTUVQWXYZ"
	letterIdBits = 6                 // 6 bits to represent a letter index
	letterIdMask = 1<<6 - 1          // all the bits set to 1
	letterMax    = 63 / letterIdBits // number of letter indices fitting in 63 bits
)

// rand function, using byte optimization, using almost all the
// 64 bits of the src.Int63() (by shifting them), in order to
// minimize the number of calls
func randString(n int) string {
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
