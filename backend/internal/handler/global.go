package handler

import (
	"context"
	"github.com/Ankr-network/kit/mlog"
	"github.com/gatewayorg/blue-dashboard/api/protos/index"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	log                = mlog.Logger("handler")
	PublicIndexHandler *PublicIndex
)

func GlobalInit() {
	PublicIndexHandler = NewPublicIndex(service.MetricsSvc)
}

func RegisterGRPC(s *grpc.Server) {
	index.RegisterPublicIndexServer(s, PublicIndexHandler)
}

func MustRegisterREST(mux *runtime.ServeMux, grpcListAddress string) {
	err := index.RegisterPublicIndexHandlerFromEndpoint(context.Background(), mux, "localhost"+grpcListAddress, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal("register public index rest error", zap.Error(err))
	}
}
