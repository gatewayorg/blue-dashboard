package host

import (
	"net"
	"time"
)

const timer = time.Second * 15

type Client struct {
	host  string
	addrs []string
}

func NewClient(host string) *Client {
	return &Client{
		host: host,
	}
}

func (c *Client) Watch() (addrsChain <-chan []string) {
	watchChain := make(chan []string)
	var err error
	go func() {
		tc := time.NewTicker(timer)
		for range tc.C {
			c.addrs, err = net.LookupHost(c.host)
			if err != nil {
				continue
			}
			watchChain <- c.addrs
		}
	}()
	return watchChain
}
