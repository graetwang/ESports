package model

import "time"

// Matches2024msi undefined
type Matches2024msi struct {
	ID         int64     `json:"id" gorm:"id"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"deleted_at"`
	ProjectId  int64     `json:"project_id" gorm:"project_id"`
	Project    string    `json:"project" gorm:"project"`
	EventId    int64     `json:"event_id" gorm:"event_id"`
	MatchEvent string    `json:"match_event" gorm:"match_event"`
	PhaseId    int64     `json:"phase_id" gorm:"phase_id"`
	Phase      string    `json:"phase" gorm:"phase"`
	GroupId    int64     `json:"group_id" gorm:"group_id"`
	Group      string    `json:"group" gorm:"group"`
	Round      string    `json:"round" gorm:"round"`
	StartTime  time.Time `json:"start_time" gorm:"start_time"`
	EndTime    time.Time `json:"end_time" gorm:"end_time"`
	Status     string    `json:"status" gorm:"status"`
	Bo         int64     `json:"bo" gorm:"bo"`
	HomeId     int64     `json:"home_id" gorm:"home_id"`
	HomeName   string    `json:"home_name" gorm:"home_name"`
	HomeLogo   string    `json:"home_logo" gorm:"home_logo"`
	HomeScore  int64     `json:"home_score" gorm:"home_score"`
	AwayId     int64     `json:"away_id" gorm:"away_id"`
	AwayName   string    `json:"away_name" gorm:"away_name"`
	AwayLogo   string    `json:"away_logo" gorm:"away_logo"`
	AwayScore  int64     `json:"away_score" gorm:"away_score"`
	Winner     string    `json:"winner" gorm:"winner"`
}

// TableName 表名称
func (*Matches2024msi) TableName() string {
	return "matches_2024MSI"
}
