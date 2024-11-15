package model

// User undefined
type User struct {
	ID       int64   `json:"id" gorm:"id"`
	Account  string  `json:"account" gorm:"account"`
	Password string  `json:"password" gorm:"password"`
	Nickname string  `json:"nickname" gorm:"nickname"`
	LiuCoin  float64 `json:"liu_coin" gorm:"liu_coin"`
}

// TableName 表名称
func (*User) TableName() string {
	return "user"
}
