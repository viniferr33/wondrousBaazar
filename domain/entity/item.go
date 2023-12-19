package entity

type Item struct {
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
	Rarity   Rarity   `json:"rarity"`
	Cost     float64  `json:"cost"`
	Desc     string   `json:"desc"`
}

type ItemRepository interface {
	Save(i *Item) error
	FindAll(itemList []Item) error
	FindAllByCategory(cat string, itemList []Item) error
	FindAllByRarity(rarity string, itemList []Item) error
	GetById(id int64, i *Item) error
}
