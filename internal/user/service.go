package user

import "strings"

type Service interface {
	GetIDByEmail(email string) (uint64, error)
	GetUsernameByID(id uint64) (string, error)
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

func (s *service) GetUsernameByID(id uint64) (string, error) {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func (s *service) SaveNewUser(email string) (uint64, error) {
	username := usernameFromEmail(email)
	usr := User{Email: email, Username: username}
	return s.repository.Insert(usr)
}

func usernameFromEmail(email string) string {
	atIndex := strings.IndexRune(email, '@')
	return email[:atIndex]
}
