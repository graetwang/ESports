package model

import "time"

// Groups undefined
type Groups struct {
	Id        int64     `json:"id" gorm:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	EventId   int64     `json:"event_id" gorm:"event_id"`
	Event     string    `json:"event" gorm:"event"`
	Phase     string    `json:"phase" gorm:"phase"`
	Name      string    `json:"name" gorm:"name"`
}

// TableName 表名称
func (*Groups) TableName() string {
	return "groups"
}
