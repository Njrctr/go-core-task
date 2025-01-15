package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	testTable := []struct {
		name  string
		mappa StringIntMap
		input struct {
			key   string
			value int
		}
		expected StringIntMap
	}{
		{
			name: "OK",
			mappa: StringIntMap{strIntMap: map[string]int{
				"key":    1,
				"qwerty": 1235123123,
			},
			},
			input: struct {
				key   string
				value int
			}{
				key:   "key2",
				value: 1,
			},
			expected: StringIntMap{strIntMap: map[string]int{
				"key":    1,
				"qwerty": 1235123123,
				"key2":   1,
			}},
		},
		{
			name: "OK key replacement",
			mappa: StringIntMap{strIntMap: map[string]int{
				"key":    1,
				"qwerty": 1235123123,
			},
			},
			input: struct {
				key   string
				value int
			}{
				key:   "key",
				value: 15,
			},
			expected: StringIntMap{strIntMap: map[string]int{
				"key":    15,
				"qwerty": 1235123123,
			}},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mappa.Add(testCase.input.key, testCase.input.value)
			equal := reflect.DeepEqual(testCase.mappa, testCase.expected)
			if !equal {
				t.Errorf("expected %v, got %v", testCase.expected, testCase.mappa)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testTable := []struct {
		name     string
		mappa    StringIntMap
		inputKey string
		expected StringIntMap
	}{
		{
			name: "OK",
			mappa: StringIntMap{strIntMap: map[string]int{
				"key":    1,
				"qwerty": 1235123123,
			},
			},
			inputKey: "qwerty",
			expected: StringIntMap{strIntMap: map[string]int{
				"key": 1,
			}},
		},
		{
			name: "OK delete not exists key",
			mappa: StringIntMap{strIntMap: map[string]int{
				"key":    1,
				"qwerty": 1235123123,
			},
			},
			inputKey: "qwerty123",
			expected: StringIntMap{strIntMap: map[string]int{
				"key":    1,
				"qwerty": 1235123123,
			}},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mappa.Delete(testCase.inputKey)
			equal := reflect.DeepEqual(testCase.mappa, testCase.expected)
			if !equal {
				t.Errorf("expected %v, got %v", testCase.expected, testCase.mappa)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	testTable := []struct {
		name  string
		mappa StringIntMap
	}{
		{
			name: "OK",
			mappa: StringIntMap{strIntMap: map[string]int{
				"key":    1,
				"qwerty": 1235123123,
			},
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.mappa.Copy()
			equal := reflect.DeepEqual(*actual, testCase.mappa)
			if !equal {
				t.Errorf("expected %v, got %v", testCase.mappa, actual)
			}
		})
	}
}

func TestExist(t *testing.T) {
	testTable := []struct {
		name      string
		mappa     StringIntMap
		searchKey string
		expected  bool
	}{
		{
			name: "True",
			mappa: StringIntMap{
				strIntMap: map[string]int{
					"key":    1,
					"qwerty": 1235123123,
				},
			},
			searchKey: "qwerty",
			expected:  true,
		},
		{
			name: "False",
			mappa: StringIntMap{
				strIntMap: map[string]int{
					"key":    1,
					"qwerty": 1235123123,
				},
			},
			searchKey: "qwerty123",
			expected:  false,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.mappa.Exist(testCase.searchKey)
			if actual != testCase.expected {
				t.Errorf("expected %v, got %v", testCase.expected, actual)
			}
		})
	}
}

func TestGet(t *testing.T) {
	testTable := []struct {
		name      string
		mappa     StringIntMap
		searchKey string
		expected  struct {
			value int
			ok    bool
		}
	}{
		{
			name: "OK1",
			mappa: StringIntMap{
				strIntMap: map[string]int{
					"key":    1,
					"qwerty": 1235123123,
				},
			},
			searchKey: "qwerty",
			expected: struct {
				value int
				ok    bool
			}{
				value: 1235123123,
				ok:    true,
			},
		},
		{
			name: "OK2",
			mappa: StringIntMap{
				strIntMap: map[string]int{
					"key":    1,
					"qwerty": 1235123123,
				},
			},
			searchKey: "qwertyqwe",
			expected: struct {
				value int
				ok    bool
			}{
				value: 0,
				ok:    false,
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual, exist := testCase.mappa.Get(testCase.searchKey)
			if exist != testCase.expected.ok || actual != testCase.expected.value {
				t.Errorf("expected %v, got %v", testCase.expected, actual)
			}
		})

	}
}
