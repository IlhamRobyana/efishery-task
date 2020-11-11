package entity

import "time"

// Timestamp is used to store the timestamps of the entities
type Timestamp struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
