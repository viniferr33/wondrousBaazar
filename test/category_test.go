package test

import (
	"encoding/json"
	"testing"
	"wondrousBaazar/domain/entity"
)

func TestCategoryUnmarshal(t *testing.T) {
	itemJson := `{"Category": "Weapon"}`

	var item entity.Item
	err := json.Unmarshal([]byte(itemJson), &item)
	if err != nil {
		t.Error(err)
	}
}
