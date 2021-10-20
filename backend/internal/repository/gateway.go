package repository

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"gorm.io/gorm"
	"time"
)

type Gateway interface {
	AddMetrics(_ context.Context, metrics *model.GatewayMetrics) error
	GetMetricsInTime(_ context.Context, start time.Time, end time.Time) ([]*model.GatewayMetrics, error)
}

type GatewayImpl struct {
	repo *gorm.DB
}

func NewGateway(db *gorm.DB) *GatewayImpl {
	return &GatewayImpl{repo: db}
}

func (g *GatewayImpl) AddMetrics(_ context.Context, metrics *model.GatewayMetrics) error {
	res := g.repo.Create(metrics)
	return res.Error
}

func (g *GatewayImpl) GetMetricsInTime(_ context.Context, start time.Time, end time.Time) ([]*model.GatewayMetrics, error) {
	var metrics []*model.GatewayMetrics
	res := g.repo.Where("created_at BETWEEN ? AND ?", start, end).Find(&metrics)
	return metrics, res.Error
}
