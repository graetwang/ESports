package model

import "time"

// Phase undefined
type Phase struct {
	ID        int64     `json:"id" gorm:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	EventId   int64     `json:"event_id" gorm:"event_id"`
	Event     string    `json:"event" gorm:"event"`
	Name      string    `json:"name" gorm:"name"`
}

// TableName 表名称
func (*Phase) TableName() string {
	return "phase"
}
