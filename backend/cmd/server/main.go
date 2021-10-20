package main

import (
	"github.com/Ankr-network/kit/app"
	"github.com/Ankr-network/kit/mlog"
	"github.com/Ankr-network/kit/rest"
	"github.com/Ankr-network/kit/rpc"
	"github.com/gatewayorg/blue-dashboard/internal/handler"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"github.com/gatewayorg/blue-dashboard/pkg/rlimit"
	"github.com/gatewayorg/blue-dashboard/share"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
	"time"
)

var log = mlog.Logger("main")

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     share.DSN,
			Usage:    "mysql dsn",
			Required: true,
		},
		&cli.StringFlag{
			Name:  share.PORT,
			Value: ":8081",
			Usage: "service port",
		},
		&cli.StringFlag{
			Name:  share.GATEWAY_SVC,
			Value: "gateway-svc",
			Usage: "service port",
		},
		&cli.StringFlag{
			Name:  share.GATEWAY_SOURCE,
			Usage: "proactive discovery of gateway ip changes (host|kubernetes)",
			Value: "kubernetes",
		},
		&cli.StringFlag{
			Name:  share.GATEWAY_NAMESPACE,
			Usage: "it is used when the gateway_source is kubernetes",
			Value: "default",
		},
		&cli.DurationFlag{
			Name:  share.SCRAPE_INTERVAL,
			Value: time.Second * 15,
			Usage: "the interval of get gateway metrics data",
		},
	}

	svr := cli.NewApp()
	svr.Action = mainServe
	svr.Flags = flags

	rlimit.SetupRLimit()

	err := svr.Run(os.Args)
	if err != nil {
		log.Fatal("Service Crash ", zap.Error(err))
	}
}

func mainServe(c *cli.Context) error {
	log.Info("init repo")
	repository.InitGlobal(c.String(share.DSN))

	log.Info("init service")
	service.GlobalInitWithConfig(&service.Config{
		GatewayService: c.String(share.GATEWAY_SVC),
		Source:         service.Source(c.String(share.GATEWAY_SOURCE)),
		Namespace:      c.String(share.GATEWAY_NAMESPACE),
		Timer:          c.Duration(share.SCRAPE_INTERVAL),
	})

	log.Info("init handler")
	handler.GlobalInit()

	log.Info("init rpc")
	rpcServer := rpc.NewServerWithConfig()
	handler.RegisterGRPC(rpcServer.Server)
	rpcServer.MustListenAndServe()

	log.Info("init rest")
	conf := rest.MustLoadConfig()
	conf.ListenAddress = c.String(share.PORT)
	restServer := rest.NewServer(conf)
	handler.MustRegisterREST(restServer.ServeMux, rpcServer.Address)
	restServer.ListenAndServed()
	app.Exit()
	return nil
}
