package test

import (
	"errors"
	"testing"
	"wondrousBaazar/domain/entity"
	"wondrousBaazar/domain/service"
	"wondrousBaazar/mock"
)

func TestItemServiceCreate(t *testing.T) {
	type TestCase struct {
		// Test Config
		testDesc string
		saveFunc func(item *entity.Item) error

		// Arrange
		name     string
		category string
		rarity   string
		cost     float64
		desc     string

		expected entity.Item
		err      error
	}

	testTable := []TestCase{
		{
			testDesc: "ShouldCreateNewItem",
			saveFunc: func(item *entity.Item) error {
				item.Id = 1
				return nil
			},
			name:     "NewItem",
			category: "Weapon",
			rarity:   "Rare",
			cost:     100.0,
			desc:     "Lorem Ipsum Dolor",

			expected: entity.Item{
				Id:       1,
				Name:     "NewItem",
				Category: "Weapon",
				Rarity:   "Rare",
				Cost:     100.0,
				Desc:     "Lorem Ipsum Dolor",
			},
			err: nil,
		},
		{
			testDesc: "ShouldNotCreateNewItem_InvalidCategory",
		},
		{
			testDesc: "ShouldNotCreateNewItem_InvalidRarity",
		},
		{
			testDesc: "ShouldNotCreateNewItem_InvalidCost",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testDesc, func(t *testing.T) {
			repository := mock.ItemRepository{
				SaveFunc: testCase.saveFunc,
			}

			s := service.NewItemService(repository)

			result, err := s.Create(
				testCase.name,
				testCase.category,
				testCase.rarity,
				testCase.cost,
				testCase.desc,
			)

			if !errors.Is(err, testCase.err) {
				t.Errorf("Received Error: %v\nExpected: %v", err, testCase.err)
			}

			if *result != testCase.expected {
				t.Errorf("Received Entity: %v\nExpected: %v", *result, testCase.expected)
			}
		})
	}
}
