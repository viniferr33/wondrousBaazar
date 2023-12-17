package service

import "wondrousBaazar/internal/domain/entity"

type ItemService struct {
	repository entity.ItemRepository
}

func NewItemService(repository entity.ItemRepository) *ItemService {
	return &ItemService{
		repository: repository,
	}
}

func (s *ItemService) Create(name string, category string, rarity string, cost float64, desc string) (*entity.Item, error) {
	var err error

	item := entity.Item{
		Name:     name,
		Category: category,
		Rarity:   rarity,
		Cost:     cost,
		Desc:     desc,
	}

	err = s.repository.Save(&item)

	return &item, err
}

func (s *ItemService) List() ([]entity.Item, error) {
	var err error
	itemList := make([]entity.Item, 0)

	err = s.repository.FindAll(itemList)

	return itemList, err
}

func (s *ItemService) GetById(id int64) (*entity.Item, error) {
	var err error
	item := entity.Item{}

	err = s.repository.GetById(id, &item)

	return &item, err
}

func (s *ItemService) ListByCategory(cat string) ([]entity.Item, error) {
	var err error
	itemList := make([]entity.Item, 0)

	err = s.repository.FindAllByCategory(cat, itemList)

	return itemList, err
}

func (s *ItemService) ListByRarity(rarity string) ([]entity.Item, error) {
	var err error
	itemList := make([]entity.Item, 0)

	err = s.repository.FindAllByRarity(rarity, itemList)

	return itemList, err
}
