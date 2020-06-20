package task

type Service interface {
	AddNewTask(task Task) error
	GetList(isCompleted bool) ([]Task, error)
	UpdateTask(task Task) error
	GetTask(id uint64) (Task, error)
	DeleteTask(id uint64) error
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

func (s *service) GetList(isCompleted bool) ([]Task, error) {
	return s.repository.FindByIsDone(isCompleted)
}

func (s *service) UpdateTask(task Task) error {
	return s.repository.Update(task)
}

func (s *service) GetTask(id uint64) (Task, error) {
	return s.repository.FindById(id)
}

func (s *service) DeleteTask(id uint64) error {
	return s.repository.Delete(id)
}
