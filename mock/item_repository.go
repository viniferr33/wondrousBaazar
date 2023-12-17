package mock

import "wondrousBaazar/domain/entity"

type ItemRepository struct {
	Items                 map[int64]entity.Item
	SaveFunc              func(*entity.Item) error
	FindAllFunc           func([]entity.Item) error
	FindAllByCategoryFunc func(string, []entity.Item) error
	FindAllByRarityFunc   func(string, []entity.Item) error
	GetByIdFunc           func(int64, *entity.Item) error
}

func (r ItemRepository) Save(i *entity.Item) error {
	if r.SaveFunc != nil {
		return r.SaveFunc(i)
	}

	return nil
}

func (r ItemRepository) FindAll(itemList []entity.Item) error {
	if r.FindAllFunc != nil {
		return r.FindAllFunc(itemList)
	}

	return nil
}

func (r ItemRepository) FindAllByCategory(cat string, itemList []entity.Item) error {
	if r.FindAllByCategoryFunc != nil {
		return r.FindAllByCategoryFunc(cat, itemList)
	}

	return nil
}

func (r ItemRepository) FindAllByRarity(rarity string, itemList []entity.Item) error {
	if r.FindAllByRarityFunc != nil {
		return r.FindAllByRarityFunc(rarity, itemList)
	}

	return nil
}

func (r ItemRepository) GetById(id int64, i *entity.Item) error {
	if r.GetByIdFunc != nil {
		return r.GetByIdFunc(id, i)
	}

	return nil
}
