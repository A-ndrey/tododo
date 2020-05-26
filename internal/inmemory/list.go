package inmemory

import (
	"fmt"
	"github.com/A-ndrey/tododo/internal/list"
)

type Sequence int64

type repository struct {
	storage map[int64]list.Item
	seq     Sequence
}

func (s *Sequence) Next() int64 {
	*s += 1
	return int64(*s)
}

func NewListRepository() list.Repository {
	return &repository{
		storage: make(map[int64]list.Item),
		seq:     0,
	}
}

func (r *repository) Insert(item list.Item) error {
	item.ID = r.seq.Next()

	if _, ok := r.storage[item.ID]; ok {
		return fmt.Errorf("can't insert: item with id=%v already exists", item.ID)
	}

	r.storage[item.ID] = item

	return nil
}

func (r *repository) Find(id int64) (list.Item, error) {
	i, ok := r.storage[id]
	if !ok {
		return i, fmt.Errorf("can't find item with id=%v", id)
	}

	return i, nil
}

func (r *repository) Update(item list.Item) error {
	if _, err := r.Find(item.ID); err != nil {
		return fmt.Errorf("can't update item: %w", err)
	}

	r.storage[item.ID] = item

	return nil
}

func (r *repository) Delete(id int64) error {
	if _, err := r.Find(id); err != nil {
		return fmt.Errorf("can't delete item: %w", err)
	}

	delete(r.storage, id)

	return nil
}
