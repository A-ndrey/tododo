package user

type Repository interface {
	FindByEmail(email string) (User, error)
	Insert(user User) (uint64, error)
}
