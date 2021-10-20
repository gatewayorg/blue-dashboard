package service

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"github.com/gatewayorg/blue-dashboard/internal/request"
	"github.com/gatewayorg/blue-dashboard/pkg/gateway_source"
	"github.com/gatewayorg/blue-dashboard/pkg/gateway_source/host"
	"github.com/gatewayorg/blue-dashboard/pkg/gateway_source/k8s"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Metrics interface {
	GetMetrics(ctx context.Context, start time.Time, end time.Time) ([]*model.GatewayMetrics, []*model.GatewayMetricsStatus, error)
	WatchMetrics(ctx context.Context)
}

type metricsImpl struct {
	sync.RWMutex
	timer          time.Duration
	gateways       map[string]request.GatewayMetrics
	watchClient    gateway_source.WatchAddrs
	gatewayRepo    repository.Gateway
	currentMetrics []*model.Metrics
}

func NewMetrics(conf *Config, repo repository.Gateway) *metricsImpl {
	var (
		client gateway_source.WatchAddrs
		err    error
	)
	if conf.Source == K8s {
		client, err = k8s.NewClient(conf.Namespace, conf.GatewayService)
		if err != nil {
			log.Fatal("k8s watch endpints", zap.String("namespace", conf.Namespace), zap.String("gateway", conf.GatewayService))
		}
	} else {
		client = host.NewClient(conf.GatewayService)
	}
	return &metricsImpl{
		gateways:    make(map[string]request.GatewayMetrics),
		timer:       conf.Timer,
		watchClient: client,
		gatewayRepo: repo,
	}
}

func (m *metricsImpl) GetMetrics(ctx context.Context, start time.Time, end time.Time) ([]*model.GatewayMetrics, []*model.GatewayMetricsStatus, error) {
	res, err := m.gatewayRepo.GetMetricsInTime(ctx, start, end)
	if err != nil {
		return nil, nil, err
	}
	currentMetrics := make([]*model.GatewayMetricsStatus, 0, len(m.gateways))
	for _, v := range m.gateways {
		m := v.GetMetrics(ctx)
		currentMetrics = append(currentMetrics, &m)
	}
	return res, currentMetrics, nil
}

func (m *metricsImpl) WatchMetrics(ctx context.Context) {
	m.watchForGatewayIpChange()
	tc := time.NewTicker(m.timer)
	for range tc.C {
		m.getGatewaysMetrics()
	}
}

func (m *metricsImpl) getGatewaysMetrics() {
	m.RLock()
	defer m.RUnlock()
	now := time.Now()
	for ip, req := range m.gateways {
		metrics, err := req.LoadMetrics(context.Background())
		if err != nil {
			log.Error("get metrics", zap.String("ip", ip), zap.Error(err))
			continue
		}

		err = m.gatewayRepo.AddMetrics(context.Background(), &model.GatewayMetrics{
			Ip:          ip,
			MetricsByte: metrics.Bytes(),
			CreatedAt:   now,
		})
		if err != nil {
			log.Error("add metrics", zap.String("ip", ip), zap.Error(err))
		}
	}

}

func (m *metricsImpl) watchForGatewayIpChange() {
	addrsChan := m.watchClient.Watch()
	go func() {
		for ips := range addrsChan {
			m.changeGateway(ips)
		}
	}()
}

func (m *metricsImpl) changeGateway(ips []string) {
	m.Lock()
	defer m.Unlock()
	gateways := make(map[string]request.GatewayMetrics)
	for _, v := range ips {
		if gwMetricsReq, ok := m.gateways[v]; ok {
			gateways[v] = gwMetricsReq
		} else {
			gateways[v] = request.NewGwMetricsImpl(v)
		}
	}
	m.gateways = gateways
}
