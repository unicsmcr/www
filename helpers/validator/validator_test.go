package validator

import (
	"strings"
	"testing"

	"github.com/hacksoc-manchester/www/helpers/rand"
)

type testCase struct {
	value  string
	result bool
}

var emailTests = []testCase{
	{rand.String(91) + "@" + rand.String(4) + "." + rand.String(3), true},
	{rand.String(50) + "@gmail.com", true},
	{rand.String(10) + "@" + rand.String(5) + "." + rand.String(3), true},
	{rand.String(20), false},
	{rand.String(99), false},
	{rand.String(50) + "@" + rand.String(10), false},
	{rand.String(10) + "." + rand.String(4), false},
	{rand.String(100), false},
}

var messageTests = []testCase{
	{rand.String(3999), true},
	{rand.String(int(rand.Src().Int63() % 400)), true},
	{strings.Repeat(" ", 50) + rand.String(1), true},
	{strings.Repeat(" ", 50) + rand.String(3950), true},
	{rand.String(4001), false},
	{strings.Repeat(" ", 100), false},
}

var nameTests = []testCase{
	{rand.String(29), true},
	{rand.String(1), true},
	{rand.String(1) + strings.Repeat(" ", 29), true},
	{rand.String(1) + strings.Repeat(" ", 30), false},
	{rand.String(31), false},
	{strings.Repeat(" ", int(rand.Src().Int63()%31)), false},
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
