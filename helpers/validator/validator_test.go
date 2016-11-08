package validator

import (
	r "github.com/hacksoc-manchester/www/helpers/rand"
	"strings"
	"testing"
)

type testCase struct {
	value  string
	result bool
}

var emailTests = []testCase{
	{r.RandString(91) + "@" + r.RandString(4) + "." + r.RandString(3), true},
	{r.RandString(50) + "@gmail.com", true},
	{r.RandString(10) + "@" + r.RandString(5) + "." + r.RandString(3), true},
	{r.RandString(20), false},
	{r.RandString(99), false},
	{r.RandString(50) + "@" + r.RandString(10), false},
	{r.RandString(10) + "." + r.RandString(4), false},
	{r.RandString(100), false},
}

var messageTests = []testCase{
	{r.RandString(3999), true},
	{r.RandString(int(r.Src().Int63() % 400)), true},
	{strings.Repeat(" ", 50) + r.RandString(1), true},
	{strings.Repeat(" ", 50) + r.RandString(3950), true},
	{r.RandString(4001), false},
	{strings.Repeat(" ", 100), false},
}

var nameTests = []testCase{
	{r.RandString(29), true},
	{r.RandString(1), true},
	{r.RandString(1) + strings.Repeat(" ", 29), true},
	{r.RandString(1) + strings.Repeat(" ", 30), false},
	{r.RandString(31), false},
	{strings.Repeat(" ", int(r.Src().Int63()%31)), false},
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
