package golibemitter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmitterOptions(t *testing.T) {
	t.Run("With empty key", func(t *testing.T) {
		eo := NewEmitterOptions("")
		assert.Equal(t, "socket.io", eo.Key)
	})
	t.Run("With key", func(t *testing.T) {
		eo := NewEmitterOptions("k")
		assert.Equal(t, "k", eo.Key)
	})
}
