package model

import "time"

// Matches undefined
type Matches struct {
	ID        int64     `json:"id" gorm:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	Name      string    `json:"name" gorm:"name"`
	ProjectId int64     `json:"project_id" gorm:"project_id"`
	StartTime time.Time `json:"start_time" gorm:"start_time"`
	EndTime   time.Time `json:"end_time" gorm:"end_time"`
	Status    string    `json:"status" gorm:"status"`
	Bo        int64     `json:"bo" gorm:"bo"`
	HomeId    int64     `json:"home_id" gorm:"home_id"`
	AwayId    int64     `json:"away_id" gorm:"away_id"`
	HomeScore int64     `json:"home_score" gorm:"home_score"`
	AwayScore int64     `json:"away_score" gorm:"away_score"`
	EventId   int64     `json:"event_id" gorm:"event_id"`
	GroupId   int64     `json:"group_id" gorm:"group_id"`
	Phase     string    `json:"phase" gorm:"phase"`
	HomeName  string    `json:"home_name" gorm:"home_name"`
	AwayName  string    `json:"away_name" gorm:"away_name"`
	Group     string    `json:"group" gorm:"group"`
}

// TableName 表名称
func (*Matches) TableName() string {
	return "matches"
}
