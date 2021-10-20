package request

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type GatewayMetrics interface {
	GetMetrics(ctx context.Context) model.GatewayMetricsStatus
	LoadMetrics(ctx context.Context) (*model.Metrics, error)
}

type gatewayMetricsImpl struct {
	ip     string
	status model.GatewayMetricsStatus
}

func NewGwMetricsImpl(ip string) *gatewayMetricsImpl {
	return &gatewayMetricsImpl{ip: ip}
}

func (p *gatewayMetricsImpl) GetMetrics(_ context.Context) model.GatewayMetricsStatus {
	return p.status
}

func (p *gatewayMetricsImpl) LoadMetrics(_ context.Context) (*model.Metrics, error) {
	status := model.GatewayMetricsStatus{
		Ip:     p.ip,
		Status: model.Unhealthy,
	}
	resp, err := http.Get(fmt.Sprintf("http://%s:9015/status/format/json", p.ip))
	if err != nil {
		p.status = status
		log.Error("get metrics", zap.String("url", p.ip), zap.Error(err))
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.status = status
		log.Error("read data of get metrics", zap.String("url", p.ip), zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()
	m := model.Metrics{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		p.status = status
		log.Error("json unmarshal of get Metrics", zap.String("url", p.ip), zap.Error(err))
		return nil, err
	}
	status.Status = model.Health
	status.Metrics = &m
	p.status = status
	return &m, nil
}
