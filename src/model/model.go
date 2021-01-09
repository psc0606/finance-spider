package model

import "time"

// 基本模型
type Model struct {
	ID        uint64     `gorm:"primary_key" json:"id" structs:"id"`
	CreatedAt time.Time  `json:"createdAt" structs:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" structs:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt" structs:"deletedAt"`
}
