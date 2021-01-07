package user

type Service interface {
	GetIDByEmail(email string) (uint64, error)
	SaveNewUser(email string) (uint64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetIDByEmail(email string) (uint64, error) {
	usr, err := s.repository.FindByEmail(email)
	if err != nil {
		return 0, err
	}

	return usr.ID, nil
}

func (s *service) SaveNewUser(email string) (uint64, error) {
	usr := User{Email: email}
	return s.repository.Insert(usr)
}
