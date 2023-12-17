package entity

type Item struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Rarity   string  `json:"rarity"`
	Cost     float64 `json:"cost"`
	Desc     string  `json:"desc"`
}

type ItemRepository interface {
	save(i *Item) error
	findAll(itemList []Item) error
	findAllByCategory(cat string, itemList []Item) error
	findAllByRarity(rarity string, itemList []Item) error
	getById(id int64, i *Item) error
}
