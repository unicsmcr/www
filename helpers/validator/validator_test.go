package validator

import (
	"github.com/hacksoc-manchester/www/helpers/rand"
	"strings"
	"testing"
)

type testCase struct {
	value  string
	result bool
}

var emailTests = []testCase{
	{rand.RandString(91) + "@" + rand.RandString(4) + "." + rand.RandString(3), true},
	{rand.RandString(50) + "@gmail.com", true},
	{rand.RandString(10) + "@" + rand.RandString(5) + "." + rand.RandString(3), true},
	{rand.RandString(20), false},
	{rand.RandString(99), false},
	{rand.RandString(50) + "@" + rand.RandString(10), false},
	{rand.RandString(10) + "." + rand.RandString(4), false},
	{rand.RandString(100), false},
}

var messageTests = []testCase{
	{rand.RandString(3999), true},
	{rand.RandString(int(rand.Src().Int63() % 400)), true},
	{strings.Repeat(" ", 50) + rand.RandString(1), true},
	{strings.Repeat(" ", 50) + rand.RandString(3950), true},
	{rand.RandString(4001), false},
	{strings.Repeat(" ", 100), false},
}

var nameTests = []testCase{
	{rand.RandString(29), true},
	{rand.RandString(1), true},
	{rand.RandString(1) + strings.Repeat(" ", 29), true},
	{rand.RandString(1) + strings.Repeat(" ", 30), false},
	{rand.RandString(31), false},
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
