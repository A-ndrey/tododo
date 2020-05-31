package list

type Repository interface {
	Insert(item Item) error
	Find(id int64) (Item, error)
	Update(item Item) error
	Delete(id int64) error
	FindByIsDone(isDone bool) ([]Item, error)
}
