package model

import "time"

type User struct {
	ID         uint64
	Username   string
	Password   string
	AliasName  string
	RoleID     uint64
	CreateTime time.Time
	Enable     bool
}
