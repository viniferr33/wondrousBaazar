package entity

import (
	"encoding/json"
	"errors"
	"strings"
)

type Category int

var ParseCategoryError error = errors.New("invalid category")

const (
	_ Category = iota
	Weapon
	Artifact
	Potion
)

func ParseCategory(category string) (Category, error) {
	switch strings.ToLower(category) {
	case "weapon":
		return Weapon, nil
	case "artifact":
		return Artifact, nil
	case "potion":
		return Potion, nil

	default:
		return 0, ParseCategoryError
	}
}

func (c *Category) UnmarshalJSON(data []byte) error {
	var categoryString string
	if err := json.Unmarshal(data, &categoryString); err != nil {
		return err
	}

	category, err := ParseCategory(categoryString)
	if err != nil {
		return err
	}

	*c = category
	return nil
}
