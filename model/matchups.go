package model

import "time"

// Matchups undefined
type Matchups struct {
	Id        int64     `json:"id" gorm:"id"`
	CreatedAt time.Time `json:"createdAt" gorm:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" gorm:"deletedAt"`
	MatchId   int64     `json:"matchId" gorm:"matchId"`
	HomeId    int64     `json:"homeId" gorm:"homeId"`
	AwayId    int64     `json:"awayId" gorm:"awayId"`
	StartTime time.Time `json:"startTime" gorm:"startTime"`
	EndTime   time.Time `json:"endTime" gorm:"endTime"`
	WinnerId  int64     `json:"winnerId" gorm:"winnerId"`
}

// TableName 表名称
func (*Matchups) TableName() string {
	return "matchups"
}
