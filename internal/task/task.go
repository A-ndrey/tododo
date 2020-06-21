package task

import (
	"time"
)

type Task struct {
	ID          uint64        `json:"id" gorm:"primary_key"`
	UserId      uint64        `json:"user_id" gorm:"not null"`
	Description string        `json:"description" gorm:"not null"`
	Duration    time.Duration `json:"duration,omitempty"`
	IsDone      bool          `json:"is_done" gorm:"type:boolean"`
	Weight      uint64        `json:"weight"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
