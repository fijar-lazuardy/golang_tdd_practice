package tdd_golang_test

import (
	"testing"

	tdd_golang "github.com/supriyadi687/tdd-golang"
)

func TestShouldReturn0GivenEmptyString(t *testing.T) {
	res, _ := tdd_golang.Add("")
	if res != 0 {
		t.Errorf("expected to be 0 got %d", res)
	}
}

func TestShouldReturnGivenNumber(t *testing.T) {
	res, _ := tdd_golang.Add("1")
	if res != 1 {
		t.Errorf("Expected to be 1 but got %d", res)
	}
}

func TestShouldReturn3GivenOneAndTwo(t *testing.T) {
	res, _ := tdd_golang.Add("1,2")
	if res != 3 {
		t.Errorf("Wrong answer, expected 3 but got %d", res)
	}
}

func TestInputShouldBeEmptyOneOrTwo(t *testing.T) {
	_, err := tdd_golang.Add("1,2,3,4,5,5")
	if err != nil {
		t.Errorf("Should not return error")
	}
}

func TestShouldAcceptCommaOrNewlineSeparatedString(t *testing.T) {
	s := "1\n2,3"
	res, _ := tdd_golang.Add(s)
	if res != 6 {
		t.Errorf("Expected 6 but got %d", res)
	}
}

// func TestShouldNotAcceptCommaAndNewlineSideBySide(t *testing.T) {
// 	s := "1,\n"
// 	res, err := tdd_golang.Add(s)
// 	if err == nil {
// 		t.Errorf("Expected error appear but got %d", res)
// 	}
// }

func TestGivingCustomDelimiter(t *testing.T) {
	s := "//;\n1;2"
	res, _ := tdd_golang.Add(s)
	if res != 3 {
		t.Errorf("Expected 3 but got %d", res)
	}
}

func TestInputCannotBeNegative(t *testing.T) {
	s := "1,-1,2"

	res, err := tdd_golang.Add(s)
	if err == nil {
		t.Errorf("Expeceted error but got %d", res)
	}
}

func TestIgnoreBigNumber(t *testing.T) {
	s := "1,1002"
	res, _ := tdd_golang.Add(s)
	if res != 1 {
		t.Errorf("Expected 1 but got %d", res)
	}
}
