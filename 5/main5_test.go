package main

import (
	"reflect"
	"testing"
)

func TestIntersectionOfSlices(t *testing.T) {
	testTable := []struct {
		name     string
		slice1   []int
		slice2   []int
		expected struct {
			hasIntersection bool
			intersection    []int
		}
	}{
		{
			name:   "OK",
			slice1: []int{1, 2, 3, 4, 5},
			slice2: []int{2, 3, 4, 5, 6},
			expected: struct {
				hasIntersection bool
				intersection    []int
			}{
				hasIntersection: true,
				intersection:    []int{2, 3, 4, 5},
			},
		},
		{
			name:   "Empty slice",
			slice1: []int{},
			slice2: []int{2, 3, 4, 5, 6},
			expected: struct {
				hasIntersection bool
				intersection    []int
			}{
				hasIntersection: false,
				intersection:    []int{},
			},
		},
		{
			name:   "all Empty slices",
			slice1: []int{},
			slice2: []int{},
			expected: struct {
				hasIntersection bool
				intersection    []int
			}{
				hasIntersection: false,
				intersection:    []int{},
			},
		},
	}
	for _, trstCase := range testTable {
		t.Run(trstCase.name, func(t *testing.T) {
			hasInter, intersection := intersectionOfSlices(trstCase.slice1, trstCase.slice2)
			if hasInter != trstCase.expected.hasIntersection {
				t.Errorf("expected %v, got %v", trstCase.expected.hasIntersection, hasInter)
			}
			if !reflect.DeepEqual(intersection, trstCase.expected.intersection) {
				t.Errorf("expected %v, got %v", trstCase.expected.intersection, intersection)
			}
		})

	}
}
