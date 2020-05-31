package list

import "time"

type Item struct {
	ID          int64         `json:"id"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	IsDone      bool          `json:"is_done"`
}
