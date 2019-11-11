package model

import "time"

type User struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string
	Gender    uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}
