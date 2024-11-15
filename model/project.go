package model

import "time"

// Project undefined
type Project struct {
	Id        int64     `json:"id" gorm:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	Name      string    `json:"name" gorm:"name"`
	Detail    string    `json:"detail" gorm:"detail"`
}

// TableName 表名称
func (*Project) TableName() string {
	return "project"
}
