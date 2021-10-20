package model

import "time"

type Role struct {
	ID         uint64
	Name       string
	Detail     string
	Enable     bool
	CreateTime time.Time
}

type Rule struct {
	ID     uint64
	Method string
	Path   string
}

type RoleRule struct {
	RoleID uint64
	RuleID uint64
}
