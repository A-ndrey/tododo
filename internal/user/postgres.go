package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where(User{Email: email}).First(&user).Error
	if err != nil {
		return User{}, fmt.Errorf("can't find user: %w", err)
	}

	return user, nil
}

func (r *repository) Insert(user User) error {
	return r.db.Save(&user).Error
}
