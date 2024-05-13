package main

import (
	"testing"
)

//go test
//go test -v
//go test -cover
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

// func TestAdd(t *testing.T) {
// 	result := add(2, 3)
// 	if result != 5 {
// 		t.Errorf("Expected 2 + 3 = 5, but result %d", result)
// 	}
// }

func TestTableAdd(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "Test1", input: []int{1, 2, 3}, expected: 6},
		{name: "Test2", input: []int{10, 10}, expected: 20},
		// {name: "Test3", input: 0, expected: 0},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.input)
			if tc.expected != result {
				t.Errorf("want: %d, got: %d", tc.expected, result)
			}
		},
		)
	}
}

// func TestMain(m *testing.M) {
// 	//m.Run(testAdd)
// 	//t.Run("TestLess", TestLess)
// }
