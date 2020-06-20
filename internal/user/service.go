package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(user User) error
	Login(user User) (uint64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Register(user User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("can't register: %w", err)
	}

	user.Password = string(hashed)

	return s.repository.Insert(user)
}

func (s *service) Login(user User) (uint64, error) {
	usr, err := s.repository.FindByEmail(user.Email)
	if err != nil {
		return 0, fmt.Errorf("can't login: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(user.Password)); err != nil {
		return 0, fmt.Errorf("can't login: %w", err)
	}

	return usr.ID, nil
}
