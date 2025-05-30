package spy

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestSpyWriter(t *testing.T) {
	spy := New()
	msg := "hello"
	n, err := spy.Write([]byte(msg))
	assert.NoError(t, err)
	assert.Equal(t, n, len(msg))
	assert.Equal(t, string(spy.Data()), msg)
}
