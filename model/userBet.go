package model

// UserBet undefined
type UserBet struct {
	ID        int64   `json:"id" gorm:"id"`
	UserId    int64   `json:"user_id" gorm:"user_id"`
	BetItemId int64   `json:"bet_item_id" gorm:"bet_item_id"`
	Amount    float64 `json:"amount" gorm:"amount"`
	IsSettled int     `json:"is_settled" gorm:"is_settled"`
}

// TableName 表名称
func (*UserBet) TableName() string {
	return "user_bet"
}
