package user

type Repository interface {
	FindByEmail(email string) (User, error)
	FindByID(id uint64) (User, error)
	Insert(user User) (uint64, error)
}
