package model

import "time"

// DbItems undefined
type DbItems struct {
	ID           int64     `json:"id" gorm:"id"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at" gorm:"deleted_at"`
	BetItemId    string    `json:"bet_item_id" gorm:"bet_item_id"`
	InitPrice    string    `json:"init_price" gorm:"init_price"`
	Price        string    `json:"price" gorm:"price"`
	MemberMaxBet string    `json:"member_max_bet" gorm:"member_max_bet"`
	ItemName     string    `json:"item_name" gorm:"item_name"`
	ItemNameEn   string    `json:"item_name_en" gorm:"item_name_en"`
	ItemNameTw   string    `json:"item_name_tw" gorm:"item_name_tw"`
	WinRate      string    `json:"win_rate" gorm:"win_rate"`
	Odds         string    `json:"odds" gorm:"odds"`
	BetListId    int64     `json:"bet_list_id" gorm:"bet_list_id"`
}

// TableName 表名称
func (*DbItems) TableName() string {
	return "db_items"
}
