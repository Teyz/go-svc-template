package entities_example_v1

import "time"

type Example struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
