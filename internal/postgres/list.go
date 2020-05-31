package postgres

import (
	"database/sql"
	"fmt"
	"github.com/A-ndrey/tododo/internal/list"
)

type repository struct {
	db *sql.DB
}

func NewListRepository(db *sql.DB) list.Repository {
	return &repository{db}
}

func (r *repository) Insert(item list.Item) error {
	query := `insert into list (description, duration, is_done) values ($1, $2, $3)`
	_, err := r.db.Exec(
		query,
		item.Description,
		item.Duration,
		item.IsDone,
	)
	if err != nil {
		return fmt.Errorf("can't insert item: %w", err)
	}

	return nil
}

func (r *repository) Find(id int64) (list.Item, error) {
	var i list.Item

	query := `select id, description, extract(epoch from duration)::integer, is_done from list where id = $1`
	err := r.db.QueryRow(query, id).Scan(&i.ID, &i.Description, &i.Duration, &i.IsDone)
	if err != nil {
		fmt.Println(err)
		return list.Item{}, fmt.Errorf("can't find item with id=%v: %w", id, err)
	}

	return i, nil
}

func (r *repository) Update(item list.Item) error {
	query := `update list set description = $1, duration = $2, is_done = $3 where id = $4`
	_, err := r.db.Exec(
		query,
		item.Description,
		item.Duration,
		item.IsDone,
		item.ID,
	)

	if err != nil {
		return fmt.Errorf("can't update item: %w", err)
	}

	return nil
}

func (r *repository) Delete(id int64) error {
	panic("implement me")
}
