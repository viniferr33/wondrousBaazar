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
