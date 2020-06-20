package task

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Insert(task Task) error {
	err := r.db.Save(&task).Error
	if err != nil {
		return fmt.Errorf("can't insert task: %w", err)
	}

	return nil
}

func (r *repository) FindById(id uint64) (Task, error) {
	var t Task

	err := r.db.First(&t, id).Error
	if err != nil {
		return Task{}, fmt.Errorf("can't find task: %w", err)
	}

	return t, nil
}

func (r *repository) Update(task Task) error {
	err := r.db.Save(&task).Error
	if err != nil {
		return fmt.Errorf("can't update task: %w", err)
	}

	return nil
}

func (r *repository) Delete(id uint64) error {
	err := r.db.Delete(Task{ID: id}).Error
	if err != nil {
		return fmt.Errorf("can't delete task: %w", err)
	}

	return nil
}

func (r *repository) FindByIsDone(isDone bool) ([]Task, error) {
	tasks := make([]Task, 0)

	err := r.db.Where(map[string]interface{}{"is_done": isDone}).Find(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("can't find by is done: %w", err)
	}

	return tasks, nil
}
