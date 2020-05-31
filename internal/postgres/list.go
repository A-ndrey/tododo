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
	query := `insert into list (description, duration, is_done, weight, created_at) values ($1, $2, $3, $4, now())`
	_, err := r.db.Exec(
		query,
		item.Description,
		item.Duration,
		item.IsDone,
		item.Weight,
	)
	if err != nil {
		return fmt.Errorf("can't insert item: %w", err)
	}

	return nil
}

func (r *repository) Find(id int64) (list.Item, error) {
	var i list.Item

	query := `
		select id, 
			   description, 
			   extract(epoch from duration)::integer, 
			   is_done,
			   weight
		from list 
		where id = $1
		  and deleted_at is null`

	err := r.db.QueryRow(query, id).Scan(
		&i.ID,
		&i.Description,
		&i.Duration,
		&i.IsDone,
		&i.Weight,
	)
	if err != nil {
		return list.Item{}, fmt.Errorf("can't find item: %w", err)
	}

	return i, nil
}

func (r *repository) Update(item list.Item) error {
	query := `
		update list 
		set description = $2,
		    duration = $3,
		    is_done = $4,
		    weight = $5,
		    updated_at = now()
		where id = $1`

	_, err := r.db.Exec(
		query,
		item.ID,
		item.Description,
		item.Duration,
		item.IsDone,
		item.Weight,
	)

	if err != nil {
		return fmt.Errorf("can't update item: %w", err)
	}

	return nil
}

func (r *repository) Delete(id int64) error {
	query := `update list set deleted_at = now() where id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("can't delete item: %w", err)
	}

	return nil
}
