package model

import "time"

// DbBetLists undefined
type DbBetLists struct {
	ID            int64     `json:"id" gorm:"id"`
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at" gorm:"deleted_at"`
	ListId        int64     `json:"list_id" gorm:"list_id"`
	BetId         string    `json:"bet_id" gorm:"bet_id"`
	Title         string    `json:"title" gorm:"title"`
	CategoryName  string    `json:"category_name" gorm:"category_name"`
	BetEndTime    string    `json:"bet_end_time" gorm:"bet_end_time"`
	BetEndTimeTxt string    `json:"bet_end_time_txt" gorm:"bet_end_time_txt"`
	Status        string    `json:"status" gorm:"status"`
	StatusTxt     string    `json:"status_txt" gorm:"status_txt"`
	TotalPrice    string    `json:"total_price" gorm:"total_price"`
	PeopleNum     string    `json:"people_num" gorm:"people_num"`
	ResultItemId  string    `json:"result_item_id" gorm:"result_item_id"`
	DynamicId     int64     `json:"dynamic_id" gorm:"dynamic_id"`
	TeamAWin      string    `json:"team_a_win" gorm:"team_a_win"`
	TeamBWin      string    `json:"team_b_win" gorm:"team_b_win"`
	MatchStatus   string    `json:"match_status" gorm:"match_status"`
	MatchBetCount string    `json:"match_bet_count" gorm:"match_bet_count"`
}

// TableName 表名称
func (*DbBetLists) TableName() string {
	return "db_bet_lists"
}
