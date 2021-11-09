package model

import "time"

type User struct {
	ID       uint64
	Username string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255)"`
	Name     string `gorm:"type:varchar(255)"`
	RoleID   uint64
	Enable   bool
	CreateAt time.Time
}

type UserSave struct {
	Name   string
	RoleID uint64
	Enable bool
}

type UserRole struct {
	ID       uint64
	Username string
	Name     string
	Enable   bool
	CreateAt time.Time

	RoleID     uint64
	RoleName   string
	RoleDetail string
}
