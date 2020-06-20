package list

import (
	"time"
)

type Item struct {
	ID          uint64        `json:"id" gorm:"primary_key"`
	Description string        `json:"description" gorm:"NOT NULL"`
	Duration    time.Duration `json:"duration,omitempty"`
	IsDone      bool          `json:"is_done" gorm:"type:boolean"`
	Weight      uint64        `json:"weight"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
