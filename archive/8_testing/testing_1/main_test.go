package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestConcatenate(t *testing.T) {
	expected := "first - second"
	result := concatenateStrings("first", "second")

	if expected != result {
		t.Errorf("Exepected %s, but result %s", expected, result)
	}
}

func TestAdd(t *testing.T) {
	result := addNumbers(5, 5)
	expected := 10

	if result != expected {
		t.Errorf("Expected %v result %v", expected, result)
	}
}

func TestConcatenateNumbers1(t *testing.T) {
	result := concatenateNumbers("1234")
	//expected := nil
	if result != nil {
		t.Errorf("result %v", result)
	}

	result = concatenateNumbers("abcd")
	fmt.Println(result)
	expected := "string does not contain only digits"
	if result.Error() != expected {
		t.Errorf("result = %v, expected = %v", result, expected)
	}

	//result := concatenateNumbers("abcd")
}

func TestConcatenateNumbers2(t *testing.T) {

	t.Run("Test1", func(t *testing.T) {
		result := concatenateNumbers("1234")
		//expected := nil
		if result != nil {
			t.Errorf("result %v", result)
		}

	})

	t.Run("Test1", func(t *testing.T) {
		result := concatenateNumbers("abcd")
		fmt.Println(result)
		expected := "string does not contain only digits"
		if result.Error() != expected {
			t.Errorf("result = %v, expected = %v", result, expected)
		}

	})
}

func TestConcatenate3(t *testing.T) {
	inputs := []struct {
		input    string
		expected error
	}{
		{"1234", nil},
		{"abcd", errors.New("string does not contain only digits")},
	}

	for _, item := range inputs {
		err := concatenateNumbers(item.input)
		if err != nil && err.Error() != item.expected.Error() {
			t.Errorf("Test failed result %v, expected %v", err, item.expected)
		}

	}
}
