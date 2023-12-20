package postgres

import (
	"context"
	"database/sql"
	"time"
	"wondrousBaazar/domain/entity"
)

type ItemRepository struct {
	db      *sql.DB
	timeout int
}

func CreateItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (r ItemRepository) Save(i *entity.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout))
	defer cancel()

	query := `
	INSERT INTO items (name, category, rarity, cost, desc) 
	VALUES ($1, $2, $3, $4, $5) RETURNING *
	`

	row, err := r.db.QueryContext(
		ctx,
		query,
		&i.Name,
		i.Category.String(),
		i.Rarity.String(),
		&i.Cost,
		&i.Desc,
	)

	if err != nil {
		return err
	}

	if err = row.Scan(&i.Id); err != nil {
		return nil
	}

	return nil
}

func (r ItemRepository) FindAll(itemList []entity.Item) error {
	//TODO implement me
	panic("implement me")
}

func (r ItemRepository) FindAllByCategory(cat string, itemList []entity.Item) error {
	//TODO implement me
	panic("implement me")
}

func (r ItemRepository) FindAllByRarity(rarity string, itemList []entity.Item) error {
	//TODO implement me
	panic("implement me")
}

func (r ItemRepository) GetById(id int64, i *entity.Item) error {
	//TODO implement me
	panic("implement me")
}
