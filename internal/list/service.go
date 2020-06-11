package list

type Service interface {
	AddNewItem(item Item) error
	GetList(isCompleted bool) ([]Item, error)
	UpdateItem(item Item) error
	GetItem(id uint) (Item, error)
	DeleteItem(id uint) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) AddNewItem(item Item) error {
	return s.repository.Insert(item)
}

func (s *service) GetList(isCompleted bool) ([]Item, error) {
	return s.repository.FindByIsDone(isCompleted)
}

func (s *service) UpdateItem(item Item) error {
	return s.repository.Update(item)
}

func (s *service) GetItem(id uint) (Item, error) {
	return s.repository.FindById(id)
}

func (s *service) DeleteItem(id uint) error {
	return s.repository.Delete(id)
}
