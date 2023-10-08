package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey"`
	NicName   string         `db:"nic_name"`
	FullName  string         `db:"full_name"`
	Email     string         `db:"email"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt *time.Time     `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at"`
}
