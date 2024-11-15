package model

import "time"

// DbLists undefined
type DbLists struct {
	ID               int64     `json:"id" gorm:"id"`
	CreatedAt        time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at" gorm:"deleted_at"`
	V                string    `json:"v" gorm:"v"`
	Key              string    `json:"key" gorm:"key"`
	NewDay           string    `json:"new_day" gorm:"new_day"`
	DateTxt          string    `json:"date_txt" gorm:"date_txt"`
	ViewType         int64     `json:"view_type" gorm:"view_type"`
	TournamentId     string    `json:"tournament_id" gorm:"tournament_id"`
	TournamentName   string    `json:"tournament_name" gorm:"tournament_name"`
	TournamentEnName string    `json:"tournament_en_name" gorm:"tournament_en_name"`
	TournamentImage  string    `json:"tournament_image" gorm:"tournament_image"`
	RoundName        string    `json:"round_name" gorm:"round_name"`
	RoundSonName     string    `json:"round_son_name" gorm:"round_son_name"`
	MatchId          string    `json:"match_id" gorm:"match_id"`
	TeamImageThumbA  string    `json:"team_image_thumb_a" gorm:"team_image_thumb_a"`
	TeamImageThumbB  string    `json:"team_image_thumb_b" gorm:"team_image_thumb_b"`
	MatchDate        string    `json:"match_date" gorm:"match_date"`
	MatchDate1       string    `json:"match_date1" gorm:"match_date1"`
	MatchTeamA       string    `json:"match_team_a" gorm:"match_team_a"`
	MatchTeamB       string    `json:"match_team_b" gorm:"match_team_b"`
	GameCount        string    `json:"game_count" gorm:"game_count"`
	StartId          string    `json:"start_id" gorm:"start_id"`
}

// TableName 表名称
func (*DbLists) TableName() string {
	return "db_lists"
}
