package service

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"time"
)

var MetricsSvc Metrics

type Source string

const (
	Host Source = "host"
	K8s  Source = "kubernetes"
)

type Config struct {
	Source
	Timer          time.Duration
	GatewayService string
	Namespace      string
}

func GlobalInitWithConfig(conf *Config) {
	MetricsSvc = NewMetrics(conf, repository.GatewayRepo)

	go MetricsSvc.WatchMetrics(context.Background())
}
