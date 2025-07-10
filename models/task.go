package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Priority    string    `json:"priority"` // Low / Medium / High
	DueDate     time.Time `json:"due_date"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
