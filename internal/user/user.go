package user

import "time"

type User struct {
	ID       uint64 `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique_index"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
