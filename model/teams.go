package model

import "time"

// Teams undefined
type Teams struct {
	Id          int64     `json:"id" gorm:"id"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"deleted_at"`
	Name        string    `json:"name" gorm:"name"`
	FoundedDate time.Time `json:"founded_date" gorm:"founded_date"`
	Region      string    `json:"region" gorm:"region"`
	Logo        string    `json:"logo" gorm:"logo"`
	ShortName   string    `json:"short_name" gorm:"short_name"`
	CreatedBy   int64     `json:"created_by" gorm:"created_by"` // 创建者
	UpdatedBy   int64     `json:"updated_by" gorm:"updated_by"` // 更新者
	DeletedBy   int64     `json:"deleted_by" gorm:"deleted_by"` // 删除者
}

// TableName 表名称
func (*Teams) TableName() string {
	return "teams"
}
