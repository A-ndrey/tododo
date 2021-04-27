package services

import (
	"fmt"
	"github.com/A-ndrey/tododo/internal/domains"
	"gorm.io/gorm"
)

type TaskService interface {
	AddNewTask(task domains.Task) (domains.Task, error)
	GetList(userId uint64) ([]domains.Task, error)
	UpdateTask(task domains.Task) error
	GetTask(userId, taskId uint64) (domains.Task, error)
	DeleteTask(userId, taskId uint64) error
}

type taskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) TaskService {
	return &taskService{db}
}

func (s *taskService) AddNewTask(task domains.Task) (domains.Task, error) {
	task.ID = 0
	err := s.db.Create(&task).Error
	if err != nil {
		return domains.Task{}, fmt.Errorf("can't insert task: %w", err)
	}

	return task, nil
}

func (s *taskService) GetList(userId uint64) ([]domains.Task, error) {
	tasks := make([]domains.Task, 0)

	err := s.db.Find(&tasks, &domains.Task{UserID: userId}).Error
	if err != nil {
		return nil, fmt.Errorf("can't find tasks: %w", err)
	}

	return tasks, nil
}

func (s *taskService) UpdateTask(task domains.Task) error {
	err := s.db.Save(&task).Error
	if err != nil {
		return fmt.Errorf("can't update task: %w", err)
	}

	return nil
}

func (s *taskService) GetTask(userId, taskId uint64) (domains.Task, error) {
	task := domains.Task{}

	err := s.db.Take(&task, &domains.Task{ID: taskId, UserID: userId}).Error
	if err != nil {
		return domains.Task{}, fmt.Errorf("can't find task: %w", err)
	}

	return task, nil
}

func (s *taskService) DeleteTask(userId, taskId uint64) error {
	err := s.db.Delete(&domains.Task{ID: taskId, UserID: userId}).Error
	if err != nil {
		return fmt.Errorf("can't delete task: %w", err)
	}

	return nil
}
