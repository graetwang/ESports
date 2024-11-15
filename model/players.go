package model

import "time"

// Players undefined
type Players struct {
	Id        int64     `json:"id" gorm:"id"`
	CreatedAt time.Time `json:"createdAt" gorm:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" gorm:"deletedAt"`
	PlayerId  int64     `json:"playerId" gorm:"playerId"`
	Name      string    `json:"name" gorm:"name"`
	Nickname  string    `json:"nickname" gorm:"nickname"`
	Role      string    `json:"role" gorm:"role"`
	TeamId    int64     `json:"teamId" gorm:"teamId"`
	Photo     string    `json:"photo" gorm:"photo"`
}

// TableName 表名称
func (*Players) TableName() string {
	return "players"
}
