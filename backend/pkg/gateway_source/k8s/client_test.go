package k8s

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClient_WatchEndpointName(t *testing.T) {
	c, err := NewProxy("127.0.0.1:8001", "default", "dccn-notifier")
	assert.NoError(t, err)
	wc := c.Watch()
	assert.NoError(t, err)
	tc := time.NewTimer(time.Second * 30)

	for {
		select {
		case res := <-wc:
			fmt.Println(res)
		case <-tc.C:
			return
		}
	}
}
