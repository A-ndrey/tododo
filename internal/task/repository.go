package task

type Repository interface {
	Insert(task Task) error
	FindById(id uint64) (Task, error)
	Update(task Task) error
	Delete(id uint64) error
	FindByIsDone(isDone bool) ([]Task, error)
}
