package model

import "time"

// Videos undefined
type Videos struct {
	ID         int64     `json:"id" gorm:"id"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"deleted_at"`
	Section    string    `json:"section" gorm:"section"`
	Title      string    `json:"title" gorm:"title"`
	Cover      string    `json:"cover" gorm:"cover"`
	Duration   string    `json:"duration" gorm:"duration"`
	ViewCounts int64     `json:"view_counts" gorm:"view_counts"`
	Likes      int64     `json:"likes" gorm:"likes"`
	Author     string    `json:"author" gorm:"author"`
	Address    string    `json:"address" gorm:"address"`
}

// TableName 表名称
func (*Videos) TableName() string {
	return "videos"
}
