package service

import (
	"errors"
	"testing"
	"wondrousBaazar/domain/entity"
	"wondrousBaazar/mock"
)

func TestItemService_Create(t *testing.T) {
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
				Category: entity.Category(1),
				Rarity:   entity.Rarity(1),
				Cost:     100.0,
				Desc:     "Lorem Ipsum Dolor",
			},
			err: nil,
		},
		{
			testDesc: "ShouldNotCreateNewItem_InvalidCategory",
			saveFunc: nil,
			name:     "InvalidCategoryItem",
			category: "invalid",
			rarity:   "Rare",
			cost:     110,
			desc:     "Lorem Ipsum",
			expected: entity.Item{},
			err:      entity.ParseCategoryError,
		},
		{
			testDesc: "ShouldNotCreateNewItem_InvalidRarity",
			saveFunc: nil,
			name:     "InvalidRarityItem",
			category: "Weapon",
			rarity:   "invalid",
			cost:     110,
			desc:     "Lorem Ipsum",
			expected: entity.Item{},
			err:      entity.ParseRarityError,
		},
		{
			testDesc: "ShouldNotCreateNewItem_InvalidCost",
			saveFunc: nil,
			name:     "InvalidCostItem",
			category: "Weapon",
			rarity:   "Rare",
			cost:     -100,
			desc:     "Lorem Ipsum",
			expected: entity.Item{},
			err:      entity.InvalidCostError,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testDesc, func(t *testing.T) {
			repository := mock.ItemRepository{
				SaveFunc: testCase.saveFunc,
			}

			s := NewItemService(repository)

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

			if result != nil {
				if *result != testCase.expected {
					t.Errorf("Received Entity: %v\nExpected: %v", *result, testCase.expected)
				}
			}
		})
	}
}
