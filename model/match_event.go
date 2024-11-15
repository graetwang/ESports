package model

import "time"

// MatchEvent undefined
type MatchEvent struct {
	Id        int64     `json:"id" gorm:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	Name      string    `json:"name" gorm:"name"`
	StartTime time.Time `json:"start_time" gorm:"start_time"`
	EndTime   time.Time `json:"end_time" gorm:"end_time"`
	Project   int64     `json:"project" gorm:"project"`
}

// TableName 表名称
func (*MatchEvent) TableName() string {
	return "match_event"
}
