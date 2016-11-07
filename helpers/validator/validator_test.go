package validator

import (
	h "github.com/hacksoc-manchester/www/helpers"
	"strings"
	"testing"
)

type testCase struct {
	value  string
	result bool
}

var emailTests = []testCase{
	{h.RandString(91) + "@" + h.RandString(4) + "." + h.RandString(3), true},
	{h.RandString(50) + "@gmail.com", true},
	{h.RandString(10) + "@" + h.RandString(5) + "." + h.RandString(3), true},
	{h.RandString(20), false},
	{h.RandString(99), false},
	{h.RandString(50) + "@" + h.RandString(10), false},
	{h.RandString(10) + "." + h.RandString(4), false},
	{h.RandString(100), false},
}

var messageTests = []testCase{
	{h.RandString(3999), true},
	{h.RandString(int(h.Src().Int63() % 400)), true},
	{strings.Repeat(" ", 50) + h.RandString(1), true},
	{strings.Repeat(" ", 50) + h.RandString(3950), true},
	{h.RandString(4001), false},
	{strings.Repeat(" ", 100), false},
}

var nameTests = []testCase{
	{h.RandString(29), true},
	{h.RandString(1), true},
	{h.RandString(1) + strings.Repeat(" ", 29), true},
	{h.RandString(1) + strings.Repeat(" ", 30), false},
	{h.RandString(31), false},
	{strings.Repeat(" ", int(h.Src().Int63()%31)), false},
}

func TestIsValidEmail(t *testing.T) {
	for _, emailTest := range emailTests {
		res := IsValidEmail(emailTest.value)
		if res != emailTest.result {
			t.Error(
				"For ", emailTest.value,
				"expected result ", emailTest.result,
				"got ", res,
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
