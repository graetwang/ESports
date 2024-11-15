package model

import "time"

// EventStandings undefined
type EventStandings struct {
	ID                     int64     `json:"id" gorm:"id"`
	CreatedAt              time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt              time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt              time.Time `json:"deleted_at" gorm:"deleted_at"`
	EventId                int64     `json:"event_id" gorm:"event_id"`
	Event                  string    `json:"event" gorm:"event"`
	TeamId                 int64     `json:"team_id" gorm:"team_id"`
	Team                   string    `json:"team" gorm:"team"`
	TeamLogo               string    `json:"team_logo" gorm:"team_logo"`
	MatchesWon             int64     `json:"matches_won" gorm:"matches_won"`
	MatchesLost            int64     `json:"matches_lost" gorm:"matches_lost"`
	DamagePerMinute        float64   `json:"damage_per_minute" gorm:"damage_per_minute"`
	GoldPerMinute          float64   `json:"gold_per_minute" gorm:"gold_per_minute"`
	CreepScorePerMinute    float64   `json:"creep_score_per_minute" gorm:"creep_score_per_minute"`
	SmallDragonControlRate float64   `json:"small_dragon_control_rate" gorm:"small_dragon_control_rate"`
	AverageKillsPerGame    float64   `json:"average_kills_per_game" gorm:"average_kills_per_game"`
	AverageAssistsPerGame  float64   `json:"average_assists_per_game" gorm:"average_assists_per_game"`
	AverageDeathsPerGame   float64   `json:"average_deaths_per_game" gorm:"average_deaths_per_game"`
	BaronControlRate       float64   `json:"baron_control_rate" gorm:"baron_control_rate"`
}

// TableName 表名称
func (*EventStandings) TableName() string {
	return "event_standings"
}
