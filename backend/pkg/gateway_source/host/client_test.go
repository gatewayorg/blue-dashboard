package host

import (
	"testing"
	"time"
)

func TestClient_Watch(t *testing.T) {
	c := NewClient("apis.ankr.com")
	ac := c.Watch()
	tc := time.NewTicker(time.Second * 30)
	for {
		select {
		case addrs := <-ac:
			t.Log(addrs)
		case <-tc.C:
			return
		}
	}
}
