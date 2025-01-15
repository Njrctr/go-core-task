package main

import "testing"

func TestStringFromAll(t *testing.T) {
	testTable := []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			name:     "OK1",
			input:    []interface{}{42, 052, 0x2A, 3.14, "Golang", true, complex64(1 + 2i)},
			expected: "4242423.14Golangtrue(1+2i)",
		},
		{
			name:     "OK2",
			input:    []interface{}{true, complex64(1 + 2i), 3.14, "Golang", 42, 052, 0x2A},
			expected: "true(1+2i)3.14Golang424242",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if actual := stringFromAll(testCase.input...); actual != testCase.expected {
				t.Errorf("expected %s, got %s", testCase.expected, actual)
			}
		})

	}
}

func TestHashWithSaltInMiddle(t *testing.T) {
	testTalse := []struct {
		name     string
		input    []rune
		expected string
	}{
		{
			name:     "OK1",
			input:    []rune("4242423.14Golangtrue(1+2i)"),
			expected: "53f2f60ac6c41389d3ed3d84d88d8c2860bf8981c677be18243a6f35a6b6a1b3",
		},
		{
			name:     "OK2",
			input:    []rune("true(1+2i)3.14Golang424242"),
			expected: "0b2cf46771eecef4618176dce1a9db2e47db67ab3a4000359afb2c6e8fee1a53",
		},
	}
	for _, testCase := range testTalse {
		t.Run(testCase.name, func(t *testing.T) {
			if actual := hashWithSaltInMiddle(testCase.input); actual != testCase.expected {
				t.Errorf("expected %s, got %s", testCase.expected, actual)
			}
		})

	}
}

func TestAllTypesToString(t *testing.T) {
	testTable := []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			name:     "OK1",
			input:    []interface{}{42, 052, 0x2A, 3.14, "Golang", true, complex64(1 + 2i)},
			expected: "int\nint\nint\nfloat64\nstring\nbool\ncomplex64",
		},
		{
			name:     "OK2",
			input:    []interface{}{true, complex64(1 + 2i), 3.14, "Golang", 42, 052, 0x2A},
			expected: "bool\ncomplex64\nfloat64\nstring\nint\nint\nint",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if actual := allTypesToString(testCase.input...); actual != testCase.expected {
				t.Errorf("expected %s, got %s", testCase.expected, actual)
			}
		})

	}
}
