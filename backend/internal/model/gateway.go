package model

import (
	"encoding/json"
	"time"
)

type ServerZone struct {
	InBytes  uint64 `json:"inBytes"`
	OutBytes uint64 `json:"outBytes"`
	//OverCounts struct {
	//	OneXx              uint64  `json:"1xx"`
	//	TwoXx              uint64  `json:"2xx"`
	//	ThreeXx            uint64  `json:"3xx"`
	//	FourXx             uint64  `json:"4xx"`
	//	FiveXx             uint64  `json:"5xx"`
	//	Bypass             uint64  `json:"bypass"`
	//	Expired            uint64  `json:"expired"`
	//	Hit                uint64  `json:"hit"`
	//	InBytes            uint64  `json:"inBytes"`
	//	MaxIntegerSize     uint64 `json:"maxIntegerSize"`
	//	Miss               uint64  `json:"miss"`
	//	OutBytes           uint64  `json:"outBytes"`
	//	RequestCounter     uint64  `json:"requestCounter"`
	//	RequestMsecCounter uint64  `json:"requestMsecCounter"`
	//	Revalidated        uint64  `json:"revalidated"`
	//	Scarce             uint64  `json:"scarce"`
	//	Stale              uint64  `json:"stale"`
	//	Updating           uint64  `json:"updating"`
	//} `json:"overCounts"`
	RequestBuckets struct {
		Counters []interface{} `json:"counters"`
		Msecs    []interface{} `json:"msecs"`
	} `json:"requestBuckets"`
	RequestCounter     int64 `json:"requestCounter"`
	RequestMsec        int64 `json:"requestMsec"`
	RequestMsecCounter int64 `json:"requestMsecCounter"`
	RequestMsecs       struct {
		Msecs []int64 `json:"msecs"`
		Times []int64 `json:"times"`
	} `json:"requestMsecs"`
	Responses struct {
		OneXx       uint64 `json:"1xx"`
		TwoXx       uint64 `json:"2xx"`
		ThreeXx     uint64 `json:"3xx"`
		FourXx      uint64 `json:"4xx"`
		FiveXx      uint64 `json:"5xx"`
		Bypass      int64  `json:"bypass"`
		Expired     int64  `json:"expired"`
		Hit         int64  `json:"hit"`
		Miss        int64  `json:"miss"`
		Revalidated int64  `json:"revalidated"`
		Scarce      int64  `json:"scarce"`
		Stale       int64  `json:"stale"`
		Updating    int64  `json:"updating"`
	} `json:"responses"`
}

type Metrics struct {
	Connections struct {
		Accepted uint64 `json:"accepted"`
		Active   uint64 `json:"active"`
		Handled  uint64 `json:"handled"`
		Reading  uint64 `json:"reading"`
		Requests uint64 `json:"requests"`
		Waiting  uint64 `json:"waiting"`
		Writing  uint64 `json:"writing"`
	} `json:"connections"`
	HostName     string                `json:"hostName"`
	LoadMsec     int64                 `json:"loadMsec"`
	NginxVersion string                `json:"nginxVersion"`
	NowMsec      int64                 `json:"nowMsec"`
	ServerZones  map[string]ServerZone `json:"serverZones"`
	SharedZones  struct {
		MaxSize  uint64 `json:"maxSize"`
		Name     string `json:"name"`
		UsedNode uint64 `json:"usedNode"`
		UsedSize uint64 `json:"usedSize"`
	} `json:"sharedZones"`
}

func (m *Metrics) Bytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}

type MetricsStatus uint8

const (
	Unhealthy MetricsStatus = iota
	Health
)

type GatewayMetricsStatus struct {
	Ip string
	*Metrics
	Status MetricsStatus
}

type GatewayMetrics struct {
	ID          uint64 `gorm:"primaryKey"`
	Ip          string `gorm:"type:varchar(50)"`
	MetricsByte []byte
	CreatedAt   time.Time
}
