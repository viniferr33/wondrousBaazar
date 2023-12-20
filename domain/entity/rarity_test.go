package entity

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestRarity_MarshalJSON(t *testing.T) {
	type TestCase struct {
		name     string
		rarity   Rarity
		expected string
		err      error
	}

	testTable := []TestCase{}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := json.Marshal(testCase.rarity)

			if !errors.Is(err, testCase.err) {
				t.Errorf("Expected error: %v\nReceived: %v", testCase.err, err)
			}

			if string(result) != testCase.expected {
				t.Errorf("Expected result: %s\nReceived: %s", testCase.expected, result)
			}
		})
	}
}

func TestRarity_UnmarshalJSON(t *testing.T) {
	type TestCase struct {
		name     string
		input    string
		expected Rarity
		err      error
	}

	testTable := []TestCase{}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var result Rarity
			err := json.Unmarshal([]byte(testCase.input), &result)

			if !errors.Is(err, testCase.err) {
				t.Errorf("Expected error: %v\nReceived: %v", testCase.err, err)
			}

			if result != testCase.expected {
				t.Errorf("Expected result: %v\nReceived: %v", testCase.expected, result)
			}
		})
	}
}

func TestRarity_String(t *testing.T) {
	type TestCase struct {
		name     string
		rarity   Rarity
		expected string
	}

	testTable := []TestCase{}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.rarity.String()

			if result != testCase.expected {
				t.Errorf("Expected result: %s\nReceived: %s", testCase.expected, result)
			}
		})
	}
}
