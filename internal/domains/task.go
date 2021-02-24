package domains

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint64 `json:"id" gorm:"primary_key"`
	UserID      uint64 `json:"user_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done" gorm:"type:boolean"`
	Weight      uint64 `json:"weight"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}