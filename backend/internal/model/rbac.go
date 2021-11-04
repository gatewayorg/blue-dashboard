package model

import (
	"github.com/gatewayorg/blue-dashboard/pkg/fields"
	"time"
)

type Role struct {
	ID       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(255)"`
	Detail   string `gorm:"type:varchar(255)"`
	Enable   bool
	RuleIDs  fields.Uint64s `gorm:"type:text"`
	CreateAt time.Time
}

type Rule struct {
	ID      uint64
	Service string `gorm:"type:varchar(255)"`
	Method  string `gorm:"type:varchar(255)"`
	Detail  string `gorm:"type:varchar(255)"`
}
