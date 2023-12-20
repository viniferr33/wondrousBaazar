package entity

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestCategory_MarshalJSON(t *testing.T) {
	type TestCase struct {
		name     string
		category Category
		expected string
		err      error
	}

	testTable := []TestCase{
		{"ShouldMarshalWeapon", Weapon, `"Weapon"`, nil},
		{"ShouldMarshalArtifact", Artifact, `"Artifact"`, nil},
		{"ShouldMarshalPotion", Potion, `"Potion"`, nil},
		{"ShouldNotMarshal_InvalidCategory", Category(42), "", ParseCategoryError},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := json.Marshal(testCase.category)

			if !errors.Is(err, testCase.err) {
				t.Errorf("Expected error: %v\nReceived: %v", testCase.err, err)
			}

			if string(result) != testCase.expected {
				t.Errorf("Expected result: %s\nReceived: %s", testCase.expected, result)
			}
		})
	}
}

func TestCategory_UnmarshalJSON(t *testing.T) {
	type TestCase struct {
		name     string
		input    string
		expected Category
		err      error
	}

	testTable := []TestCase{
		{"ShouldUnmarshalWeapon", `"Weapon"`, Weapon, nil},
		{"ShouldUnmarshalArtifact", `"Artifact"`, Artifact, nil},
		{"ShouldUnmarshalPotion", `"Potion"`, Potion, nil},
		{"ShouldNotUnmarshal_InvalidCategory", `"InvalidCategory"`, Category(0), ParseCategoryError},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var result Category
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
		category Category
		expected string
	}

	testTable := []TestCase{
		{"ShouldStringWeapon", Weapon, "Weapon"},
		{"ShouldStringArtifact", Artifact, "Artifact"},
		{"ShouldStringPotion", Potion, "Potion"},
		{"ShouldNotString_InvalidCategory", Category(42), "Invalid"},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.category.String()

			if result != testCase.expected {
				t.Errorf("Expected result: %s\nReceived: %s", testCase.expected, result)
			}
		})
	}
}
