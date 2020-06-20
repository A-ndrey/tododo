package list

type Repository interface {
	Insert(item Item) error
	FindById(id uint64) (Item, error)
	Update(item Item) error
	Delete(id uint64) error
	FindByIsDone(isDone bool) ([]Item, error)
}
