package repository

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"time"
)

type Gateway interface {
	AddMetrics(_ context.Context, metrics *model.GatewayMetrics) error
	GetMetricsInTime(_ context.Context, start time.Time, end time.Time) ([]*model.GatewayMetrics, error)
}

type gatewayImpl struct {
	*DB
}

func NewGateway(db *DB) *gatewayImpl {
	return &gatewayImpl{DB: db}
}

func (g *gatewayImpl) AddMetrics(ctx context.Context, metrics *model.GatewayMetrics) error {
	res := g.GetTxFromContext(ctx).Create(metrics)
	return res.Error
}

func (g *gatewayImpl) GetMetricsInTime(ctx context.Context, start time.Time, end time.Time) ([]*model.GatewayMetrics, error) {
	var metrics []*model.GatewayMetrics
	res := g.GetTxFromContext(ctx).Where("created_at BETWEEN ? AND ?", start, end).Find(&metrics)
	return metrics, res.Error
}
