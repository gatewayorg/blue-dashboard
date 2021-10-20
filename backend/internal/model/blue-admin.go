package model

import (
	"encoding/json"
)

type GetUri struct {
	PassWd   string `json:"passwd"`
	URI      string `json:"uri"`
	Username string `json:"username"`
}

func (g *GetUri) Bytes() (res []byte) {
	res, _ = json.Marshal(g)
	return
}

type UriInfo struct {
	PassWd   string `json:"passwd"`
	SrvUri   string `json:"srvuri"`
	URI      string `json:"uri"`
	Username string `json:"username"`
}

func (g *UriInfo) Bytes() (res []byte) {
	res, _ = json.Marshal(g)
	return
}

type UriRes struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	URI     string `json:"uri"`
}
