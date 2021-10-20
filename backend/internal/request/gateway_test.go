package request

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGatewayMetricsImpl_GetMetrics(t *testing.T) {
	m := NewGwMetricsImpl("127.0.0.1")
	metrics, err := m.LoadMetrics(context.Background())
	assert.NoError(t, err)
	t.Log(metrics)
}
