package model

import "time"

// GroupStandings undefined
type GroupStandings struct {
	Id            int64     `json:"id" gorm:"id"`
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at" gorm:"deleted_at"`
	GroupId       int64     `json:"group_id" gorm:"group_id"`
	TeamId        int64     `json:"team_id" gorm:"team_id"`
	MatchesPlayed int64     `json:"matches_played" gorm:"matches_played"`
	MatchesWon    int64     `json:"matches_won" gorm:"matches_won"`
	MatchesLost   int64     `json:"matches_lost" gorm:"matches_lost"`
	Points        int64     `json:"points" gorm:"points"`
}

// TableName 表名称
func (*GroupStandings) TableName() string {
	return "group_standings"
}
