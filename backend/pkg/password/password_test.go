package password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	hashed, err := Hash("test")
	assert.NoError(t, err)
	t.Log(hashed)
	assert.NoError(t, Verify(hashed, "test"))
	assert.Error(t, Verify(hashed, "test1"))
}
