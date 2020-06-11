package list

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Insert(item Item) error {
	err := r.db.Save(&item).Error
	if err != nil {
		return fmt.Errorf("can't insert item: %w", err)
	}

	return nil
}

func (r *repository) FindById(id uint) (Item, error) {
	var i Item

	err := r.db.First(&i, id).Error
	if err != nil {
		return Item{}, fmt.Errorf("can't find item: %w", err)
	}

	return i, nil
}

func (r *repository) Update(item Item) error {
	err := r.db.Save(&item).Error
	if err != nil {
		return fmt.Errorf("can't update item: %w", err)
	}

	return nil
}

func (r *repository) Delete(id uint) error {
	err := r.db.Delete(Item{ID: id}).Error
	if err != nil {
		return fmt.Errorf("can't delete item: %w", err)
	}

	return nil
}

func (r *repository) FindByIsDone(isDone bool) ([]Item, error) {
	items := make([]Item, 0)

	err := r.db.Where(map[string]interface{}{"is_done": isDone}).Find(&items).Error
	if err != nil {
		return nil, fmt.Errorf("can't find by is done: %w", err)
	}

	return items, nil
}
