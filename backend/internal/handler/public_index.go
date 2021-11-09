package handler

import (
	"context"
	"encoding/json"
	"github.com/gatewayorg/blue-dashboard/api/protos/index"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PublicIndex struct {
	gatewaySvc service.Metrics
}

func NewPublicIndex(gatewaySvc service.Metrics) *PublicIndex {
	return &PublicIndex{
		gatewaySvc: gatewaySvc,
	}
}

func (p *PublicIndex) Index(ctx context.Context, req *index.IndexReq) (*index.IndexResp, error) {
	gatewayMetrics, gatewayMetricsStatus, err := p.gatewaySvc.GetMetrics(ctx, time.Unix(int64(req.Start), 0), time.Unix(int64(req.End), 0))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	//log.Info("gatewatMetrics", zap.Any("data", gatewayMetrics))
	//log.Info("gatewayMetricsStatus", zap.Any("data", gatewayMetricsStatus))

	return &index.IndexResp{Data: gatewayMetricsToPb(gatewayMetrics, gatewayMetricsStatus)}, nil
}

func gatewayMetricsToPb(data []*model.GatewayMetrics, status []*model.GatewayMetricsStatus) []*index.GatewayInfo {
	sharedMems := make(map[string][]*index.SharedMem)
	connections := make(map[string][]*index.Connections)
	for _, metrics := range data {
		t := timestamppb.New(metrics.CreatedAt)

		var m model.Metrics
		err := json.Unmarshal(metrics.MetricsByte, &m)
		if err != nil {
			continue
		}
		_, ok := sharedMems[metrics.Ip]
		if !ok {
			sharedMems[metrics.Ip] = make([]*index.SharedMem, 0, len(data)/3)
		}
		sharedMems[metrics.Ip] = append(sharedMems[metrics.Ip], &index.SharedMem{
			MaxSize:  m.SharedZones.MaxSize,
			UsedSize: m.SharedZones.UsedSize,
			UsedNode: m.SharedZones.UsedSize,
			Time:     t,
		})

		_, ok = connections[metrics.Ip]
		if !ok {
			connections[metrics.Ip] = make([]*index.Connections, 0, len(data)/3)
		}
		connections[metrics.Ip] = append(connections[metrics.Ip], &index.Connections{
			Accepted: m.Connections.Accepted,
			Active:   m.Connections.Active,
			Handled:  m.Connections.Handled,
			Reading:  m.Connections.Reading,
			Requests: m.Connections.Requests,
			Waiting:  m.Connections.Waiting,
			Writing:  m.Connections.Writing,
			Time:     t,
		})
	}
	res := make([]*index.GatewayInfo, 0, len(status))
	for _, info := range status {
		reqTotal := &index.RequestTotal{}
		cacheTotal := &index.CacheTotal{}
		if info.Metrics != nil {
			for _, serverZones := range info.ServerZones {
				reqTotal.InBytes += serverZones.InBytes
				reqTotal.OutBytes += serverZones.OutBytes
				reqTotal.X1Xx += serverZones.Responses.OneXx
				reqTotal.X2Xx += serverZones.Responses.TwoXx
				reqTotal.X3Xx += serverZones.Responses.ThreeXx
				reqTotal.X4Xx += serverZones.Responses.FourXx
				reqTotal.X5Xx += serverZones.Responses.FiveXx

				cacheTotal.Miss += serverZones.Responses.Miss
				cacheTotal.Bypass += serverZones.Responses.Bypass
				cacheTotal.Expired += serverZones.Responses.Expired
				cacheTotal.Stale += serverZones.Responses.Stale
				cacheTotal.Updating += serverZones.Responses.Updating
				cacheTotal.Revalidated += serverZones.Responses.Revalidated
				cacheTotal.Hit += serverZones.Responses.Hit
				cacheTotal.Scarce += serverZones.Responses.Scarce
			}
		}

		res = append(res, &index.GatewayInfo{
			HostName:     info.HostName,
			Version:      info.NginxVersion,
			Status:       index.GatewayStatus(info.Status),
			Ip:           info.Ip,
			SharedMemory: sharedMems[info.Ip],
			Connections:  connections[info.Ip],
			RequestTotal: reqTotal,
			CacheTotal:   cacheTotal,
		})
	}
	return res
}
