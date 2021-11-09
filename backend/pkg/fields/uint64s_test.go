package fields

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUint64s_Value(t *testing.T) {
	var (
		u    = Uint64s{1, 2, 3, 3}
		newU = Uint64s{}
	)

	v, err := u.Value()
	assert.NoError(t, err)
	t.Log(v)

	err = newU.Scan([]byte("2,3,4,4"))
	assert.NoError(t, err)
	t.Log(newU)
}
