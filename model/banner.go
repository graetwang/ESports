package model

import "time"

// Banner undefined
type Banner struct {
	Id        int64     `json:"id" gorm:"id"`
	CreatedAt time.Time `json:"createdAt" gorm:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" gorm:"deletedAt"`
	Picture   string    `json:"picture" gorm:"picture"`
	Title     string    `json:"title" gorm:"title"`
}

// TableName 表名称
func (*Banner) TableName() string {
	return "banner"
}
