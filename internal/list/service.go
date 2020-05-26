package list

type Service interface {
	AddNewItem(item Item) error
	GetActualList() ([]Item, error)
	UpdateItem(item Item) error
	GetItem(id int64) (Item, error)
	DeleteItem(id int64) error
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

func (s *service) GetActualList() ([]Item, error) {
	panic("implement me")
}

func (s *service) UpdateItem(item Item) error {
	return s.repository.Update(item)
}

func (s *service) GetItem(id int64) (Item, error) {
	return s.repository.Find(id)
}

func (s *service) DeleteItem(id int64) error {
	return s.repository.Delete(id)
}
