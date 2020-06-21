package task

type Repository interface {
	Insert(task Task) error
	FindById(userId, taskId uint64) (Task, error)
	Update(task Task) error
	Delete(userId, taskId uint64) error
	FindByIsDone(userId uint64, isDone bool) ([]Task, error)
}
