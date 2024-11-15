package model


import "time"

// Background undefined
type Background struct {
	ID        int64     `json:"id" gorm:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	Picture   string    `json:"picture" gorm:"picture"`
	Color     string    `json:"color" gorm:"color"`
}

// TableName 表名称
func (*Background) TableName() string {
	return "background"
}
