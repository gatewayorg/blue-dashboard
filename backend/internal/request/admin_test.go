package request

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdminImpl_GetUriList(t *testing.T) {
	a := NewAdmin("http://127.0.0.1:5000", "ece5513c-281b-11ec-aa10-ebac4e67095b")

	err := a.CreateUri(context.Background(), &model.UriInfo{
		PassWd:   "aaa",
		SrvUri:   "http://127.0.0.1:333/jmz",
		URI:      "/jmz",
		Username: "bbb",
	})
	assert.NoError(t, err)

	ui, err := a.GetUri(context.Background(), &model.GetUri{
		URI:      "/jmz",
		PassWd:   "aaa",
		Username: "bbb",
	})
	assert.NoError(t, err)
	t.Log(ui)

}
