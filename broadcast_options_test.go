package golibemitter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBroadcastOptions(t *testing.T) {
	t.Run("With empty key and namespace", func(t *testing.T) {
		bo := NewBroadcastOptions("", "")
		assert.Equal(t, "/", bo.Namespace)
		assert.Equal(t, "socket.io#/#", bo.BroadcastChannel)
		assert.Equal(t, "socket.io-request#/#", bo.RequestChannel)
	})
	t.Run("With empty key and not empty namespace", func(t *testing.T) {
		bo := NewBroadcastOptions("", "ns")
		assert.Equal(t, "ns", bo.Namespace)
		assert.Equal(t, "socket.io#ns#", bo.BroadcastChannel)
		assert.Equal(t, "socket.io-request#ns#", bo.RequestChannel)
	})
	t.Run("With not empty key and empty namespace", func(t *testing.T) {
		bo := NewBroadcastOptions("k", "")
		assert.Equal(t, "/", bo.Namespace)
		assert.Equal(t, "k#/#", bo.BroadcastChannel)
		assert.Equal(t, "k-request#/#", bo.RequestChannel)
	})
	t.Run("With not empty key and namespace", func(t *testing.T) {
		bo := NewBroadcastOptions("k", "ns")
		assert.Equal(t, "ns", bo.Namespace)
		assert.Equal(t, "k#ns#", bo.BroadcastChannel)
		assert.Equal(t, "k-request#ns#", bo.RequestChannel)
	})
}
