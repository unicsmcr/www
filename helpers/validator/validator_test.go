package validator

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

type testCase struct {
	value  string
	result bool
}

var emailTests = []testCase{
	{randString(91) + "@" + randString(4) + "." + randString(3), true},
	{randString(50) + "@gmail.com", true},
	{randString(10) + "@" + randString(5) + "." + randString(3), true},
	{randString(20), false},
	{randString(99), false},
	{randString(50) + "@" + randString(10), false},
	{randString(10) + "." + randString(4), false},
	{randString(100), false},
}

var messageTests = []testCase{
	{randString(3999), true},
	{randString(int(src.Int63() % 400)), true},
	{strings.Repeat(" ", 50) + randString(1), true},
	{strings.Repeat(" ", 50) + randString(3950), true},
	{randString(4001), false},
	{strings.Repeat(" ", 100), false},
}

var nameTests = []testCase{
	{randString(29), true},
	{randString(1), true},
	{randString(1) + strings.Repeat(" ", 29), true},
	{randString(1) + strings.Repeat(" ", 30), false},
	{randString(31), false},
	{strings.Repeat(" ", int(src.Int63()%31)), false},
}

func TestIsValidEmail(t *testing.T) {
	for _, emailTest := range emailTests {
		res := IsValidEmail(emailTest.value)
		if res != emailTest.result {
			t.Error(
				"For", emailTest.value,
				"expected result", emailTest.result,
				"got", res,
			)
		}
	}
}

func TestIsValidMessage(t *testing.T) {
	for _, messageTest := range messageTests {
		res := IsValidMessage(messageTest.value)
		if res != messageTest.result {
			t.Error(
				"For", messageTest.value,
				"expected result", messageTest.result,
				"got", res,
			)
		}
	}
}

func TestIsValidName(t *testing.T) {
	for _, nameTest := range nameTests {
		res := IsValidName(nameTest.value)
		if res != nameTest.result {
			t.Error(
				"For", nameTest.value,
				"expected result", nameTest.result,
				"got", res,
			)
		}
	}
}

// new source for rand func, seeding it with time.Now().UnixNano()
// for different outputs at different runs
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
