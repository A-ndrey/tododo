package domains

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint64 `json:"id" gorm:"primary_key"`
	UserID      uint64 `json:"userID" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone" gorm:"type:boolean"`
	Weight      uint64 `json:"weight"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
