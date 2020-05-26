package list

import "time"

type Item struct {
	ID       int64         `json:"id"`
	Text     string        `json:"text"`
	Duration time.Duration `json:"duration"`
	IsDone   bool          `json:"is_done"`
}
