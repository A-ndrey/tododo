package task

type Service interface {
	AddNewTask(task Task) error
	GetList(userId uint64, isCompleted bool) ([]Task, error)
	UpdateTask(task Task) error
	GetTask(userId, taskId uint64) (Task, error)
	DeleteTask(userId, taskId uint64) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) AddNewTask(task Task) error {
	return s.repository.Insert(task)
}

func (s *service) GetList(userId uint64, isCompleted bool) ([]Task, error) {
	return s.repository.FindByIsDone(userId, isCompleted)
}

func (s *service) UpdateTask(task Task) error {
	return s.repository.Update(task)
}

func (s *service) GetTask(userId, taskId uint64) (Task, error) {
	return s.repository.FindById(userId, taskId)
}

func (s *service) DeleteTask(userId, taskId uint64) error {
	return s.repository.Delete(userId, taskId)
}
