package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint64 `json:"id,omitempty" gorm:"primary_key"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty" gorm:"unique_index"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
