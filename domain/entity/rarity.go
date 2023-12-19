package entity

import (
	"encoding/json"
	"errors"
	"strings"
)

type Rarity int

const (
	_ Rarity = iota
	Rare
	Epic
	Legendary
	Wondrous
)

func ParseRarity(rarity string) (Rarity, error) {
	switch strings.ToLower(rarity) {
	case "rare":
		return Rare, nil
	case "epic":
		return Epic, nil
	case "legendary":
		return Legendary, nil
	case "wondrous":
		return Wondrous, nil
	default:
		return 0, errors.New("invalid rarity")
	}
}

func (r *Rarity) UnmarshalJSON(data []byte) error {
	var rarityString string
	if err := json.Unmarshal(data, &rarityString); err != nil {
		return err
	}

	rarity, err := ParseRarity(rarityString)
	if err != nil {
		return err
	}

	*r = rarity
	return nil
}
