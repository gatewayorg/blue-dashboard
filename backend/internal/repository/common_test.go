package repository

import (
	"context"
	"fmt"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	"testing"
)

var (
	testRepo    *gorm.DB
	dsn         = "root:123@tcp(127.0.0.1:3306)/dashboard?charset=utf8mb4&parseTime=True"
	testGateway Gateway
	testMetrics = `{"hostName":"blue-gw","nginxVersion":"1.19.9","loadMsec":1634115435787,"nowMsec":1634711689546,"connections":{"active":1,"reading":0,"writing":1,"waiting":0,"accepted":24,"handled":24,"requests":995},"sharedZones":{"name":"ngx_http_vhost_traffic_status","maxSize":1048575,"usedSize":3510,"usedNode":1},"serverZones":{"_":{"requestCounter":994,"inBytes":256388,"outBytes":4025577,"responses":{"1xx":0,"2xx":993,"3xx":0,"4xx":1,"5xx":0,"miss":0,"bypass":0,"expired":0,"stale":0,"updating":0,"revalidated":0,"hit":0,"scarce":0},"requestMsecCounter":0,"requestMsec":0,"requestMsecs":{"times":[1634191715813,1634191718787,1634191721785,1634191724787,1634191727797,1634191730801,1634191733800,1634191736803,1634191739807,1634191742811,1634191745829,1634191748783,1634191751786,1634191754790,1634191757792,1634191760796,1634191763801,1634191766804,1634191769806,1634191772809,1634191775814,1634191778785,1634191781787,1634191784788,1634191787810,1634191790801,1634191793803,1634191796804,1634191799821,1634191802812,1634191805818,1634191808792,1634191811790,1634191814792,1634191817795,1634191820799,1634191823801,1634191826807,1634191829816,1634191832810,1634191835818,1634191838786,1634191841788,1634191844789,1634191847792,1634191850799,1634191853803,1634191856803,1634191859813,1634191862815,1634191865814,1634191868789,1634191871786,1634191874793,1634191877797,1634191880802,1634191883807,1634191886806,1634524510701,1634539785372,1634611768825,1634611819198,1634611834134],"msecs":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]},"requestBuckets":{"msecs":[],"counters":[]},"overCounts":{"maxIntegerSize":18446744073709551615,"requestCounter":0,"inBytes":0,"outBytes":0,"1xx":0,"2xx":0,"3xx":0,"4xx":0,"5xx":0,"miss":0,"bypass":0,"expired":0,"stale":0,"updating":0,"revalidated":0,"hit":0,"scarce":0,"requestMsecCounter":0}},"*":{"requestCounter":994,"inBytes":256388,"outBytes":4025577,"responses":{"1xx":0,"2xx":993,"3xx":0,"4xx":1,"5xx":0,"miss":0,"bypass":0,"expired":0,"stale":0,"updating":0,"revalidated":0,"hit":0,"scarce":0},"requestMsecCounter":0,"requestMsec":0,"requestMsecs":{"times":[1634191805818,1634191808792,1634191811790,1634191814792,1634191817795,1634191820799,1634191823801,1634191826807,1634191829816,1634191832810,1634191835818,1634191838786,1634191841788,1634191844789,1634191847792,1634191850799,1634191853803,1634191856803,1634191859813,1634191862815,1634191865814,1634191868789,1634191871786,1634191874793,1634191877797,1634191880802,1634191883807,1634191886806,1634524510701,1634539785372,1634611768825,1634611819198,1634611834134,1634191712810,1634191715813,1634191718787,1634191721785,1634191724787,1634191727797,1634191730801,1634191733800,1634191736803,1634191739807,1634191742811,1634191745829,1634191748783,1634191751786,1634191754790,1634191757792,1634191760796,1634191763801,1634191766804,1634191769806,1634191772809,1634191775814,1634191778785,1634191781787,1634191784788,1634191787810,1634191790801,1634191793803,1634191796804,1634191799821],"msecs":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]},"requestBuckets":{"msecs":[],"counters":[]},"overCounts":{"maxIntegerSize":18446744073709551615,"requestCounter":0,"inBytes":0,"outBytes":0,"1xx":0,"2xx":0,"3xx":0,"4xx":0,"5xx":0,"miss":0,"bypass":0,"expired":0,"stale":0,"updating":0,"revalidated":0,"hit":0,"scarce":0,"requestMsecCounter":0}}}}`
)

func TestMain(m *testing.M) {
	var err error
	testRepo, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	testGateway = NewGateway(testRepo)
	if err == nil {
		m.Run()
	} else {
		fmt.Println(err)
	}
}

func TestAutoMigrate(t *testing.T) {
	testRepo.AutoMigrate(&model.GatewayMetrics{})
}

func TestGatewayImpl_AddMetrics(t *testing.T) {
	testGateway.AddMetrics(context.Background(), &model.GatewayMetrics{
		Ip:          "191.168.0.2",
		MetricsByte: []byte(testMetrics),
		CreatedAt:   time.Now(),
	})
}

func TestGatewayImpl_GetMetricsInTime(t *testing.T) {
	metrics, err := testGateway.GetMetricsInTime(context.Background(), time.Now().Add(-time.Hour*24), time.Now())
	assert.NoError(t, err)
	t.Log(metrics)
}
