package list

type Repository interface {
	Insert(item Item) error
	FindById(id uint) (Item, error)
	Update(item Item) error
	Delete(id uint) error
	FindByIsDone(isDone bool) ([]Item, error)
}
