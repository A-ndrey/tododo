package services

import (
	"fmt"
	"github.com/A-ndrey/tododo/internal/domains"
	"gorm.io/gorm"
	"strings"
)

type UserService interface {
	GetIDByEmail(email string) (uint64, error)
	GetUsernameByID(id uint64) (string, error)
	SaveNewUser(email string) (uint64, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db}
}

func (s *userService) GetIDByEmail(email string) (uint64, error) {
	user := domains.User{Email: email}

	err := s.db.First(&user).Error
	if err != nil {
		return 0, fmt.Errorf("can't find user: %w", err)
	}

	return user.ID, nil
}

func (s *userService) GetUsernameByID(id uint64) (string, error) {
	user := domains.User{ID: id}

	err := s.db.First(&user).Error
	if err != nil {
		return "", fmt.Errorf("can't find user: %w", err)
	}

	return user.Username, nil
}

func (s *userService) SaveNewUser(email string) (uint64, error) {
	username := usernameFromEmail(email)
	user := domains.User{Email: email, Username: username}
	res := s.db.Save(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}

func usernameFromEmail(email string) string {
	atIndex := strings.IndexRune(email, '@')
	return email[:atIndex]
}
