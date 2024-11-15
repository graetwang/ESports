package model

import "time"

// Schedules undefined
type Schedules struct {
	Id         int64     `json:"id" gorm:"id"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"deleted_at"`
	MatchId    int64     `json:"match_id" gorm:"match_id"`
	MatchTime  time.Time `json:"match_time" gorm:"match_time"`
	Round      string    `json:"round" gorm:"round"`
	Weight     int64     `json:"weight" gorm:"weight"`
	MatchEvent string    `json:"match_event" gorm:"match_event"`
	Project    string    `json:"project" gorm:"project"`
}

// TableName 表名称
func (*Schedules) TableName() string {
	return "schedules"
}
