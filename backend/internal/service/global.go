package service

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"time"
)

var (
	MetricsSvc Metrics
	RbacSvc    Rbac
	UserSvc    User
)

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
	RbacSvc = NewRbac(repository.RbacRepo, repository.UserRepo)
	UserSvc = NewUser(repository.UserRepo, repository.RbacRepo)

	go MetricsSvc.WatchMetrics(context.Background())
}
