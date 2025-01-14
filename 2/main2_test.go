package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceExample(t *testing.T) {
	testTable := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "OK1",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "OK2",
			input:    []int{600, 214, 572, 921, 35, 92, 745, 725, 620, 224},
			expected: []int{600, 214, 572, 92, 620, 224},
		},
		{
			name:     "OK3",
			input:    []int{13, 21, 45, 11, 955, 3233, 109045, 1, 7, 887},
			expected: []int{},
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if actual := sliceExample(testCase.input); reflect.DeepEqual(actual, testCase.expected) != true {
				t.Errorf("SliceExample(%v) = %v, want %v", testCase.input, actual, testCase.expected)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	testTable := []struct {
		name         string
		inputSlice   []int
		inputElement int
		expected     []int
	}{
		{
			name:         "OK1",
			inputSlice:   []int{2, 4, 6, 8, 10},
			inputElement: 101,
			expected:     []int{2, 4, 6, 8, 10, 101},
		},
		{
			name:         "OK2",
			inputSlice:   []int{600, 214, 572, 921, 35, 92, 745, 725, 620, 224},
			inputElement: 101,
			expected:     []int{600, 214, 572, 921, 35, 92, 745, 725, 620, 224, 101},
		},
		{
			name:         "OK3",
			inputSlice:   []int{},
			inputElement: 101,
			expected:     []int{101},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if actual := addElements(
				testCase.inputSlice,
				testCase.inputElement); reflect.DeepEqual(actual, testCase.expected) != true {
				t.Errorf(
					"addElements(%v, %v) = %v, want %v",
					testCase.inputSlice, testCase.inputElement,
					actual,
					testCase.expected,
				)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	testTable := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "OK1",
			input:    []int{2, 4, 6, 8, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "OK2",
			input:    []int{},
			expected: []int{},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual := copySlice(testCase.input)
			if reflect.DeepEqual(actual, testCase.expected) != true {
				t.Errorf("copySlice(%v) = %v, want %v", testCase.input, actual, testCase.expected)
			} else if len(actual) > 0 && unsafe.Pointer(&actual[0]) == unsafe.Pointer(&testCase.input[0]) {
				t.Errorf("adress &actual[0] = %v, and &testCase.input[0] = %v", unsafe.Pointer(&actual[0]), unsafe.Pointer(&testCase.input[0]))
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	testTable := []struct {
		name              string
		inputSlice        []int
		inputElementIndex int
		expected          []int
		wantedErrStr      string
	}{
		{
			name:              "OK1",
			inputSlice:        []int{2, 4, 6, 8, 10},
			inputElementIndex: 2,
			expected:          []int{2, 4, 8, 10},
		},
		{
			name:              "OK2",
			inputSlice:        []int{600, 214, 572, 921, 35, 92, 745, 725, 620, 224},
			inputElementIndex: 2,
			expected:          []int{600, 214, 921, 35, 92, 745, 725, 620, 224},
		},
		{
			name:              "Index out of range",
			inputSlice:        []int{600, 214, 572},
			inputElementIndex: 3,
			wantedErrStr:      "index out of range",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := removeElement(testCase.inputSlice, testCase.inputElementIndex)
			if err != nil {
				fmt.Println(err.Error() == testCase.wantedErrStr)
				if err.Error() == testCase.wantedErrStr {
					return
				}
				t.Errorf("removeElement(%v, %v) return error %v", testCase.inputSlice, testCase.inputElementIndex, err.Error())
			}
			if reflect.DeepEqual(actual, testCase.expected) != true {
				t.Errorf("removeElement(%v, %v) = %v, want %v", testCase.inputSlice, testCase.inputElementIndex, actual, testCase.expected)
			}
		})

	}
}
