package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type Task struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Description    string    `json:"description" db:"description"`
	IsDone         bool      `json:"is_done" db:"is_done"`
	CompletionDate string    `json:"completion_date" db:"completion_date"`
	RequestedBy    string    `json:"requested_by" db:"requested_by"`
	ExecutedBy     string    `json:"executed_by" db:"executed_by"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type Tasks []Task

func (t Task) TableName() string {
	return "task"
}

func (t Task) String() string {
	jp, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(jp)
}

func (t Tasks) String() string {
	jp, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(jp)
}
