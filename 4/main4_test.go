package main

import (
	"reflect"
	"testing"
)

func TestGetUniqueBySlices(t *testing.T) {
	testTable := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "No one unique",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "all is unique",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"d", "e", "f"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "OK",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"d", "e", "f", "c"},
			expected: []string{"a", "b"},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual := getUniqueBySlices(testCase.slice1, testCase.slice2)
			if !equalSlices(actual, testCase.expected) {
				t.Errorf("expected %v, actual %v", testCase.expected, actual)
			}
		})
	}
}

func equalSlices(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}
