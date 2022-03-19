package golibemitter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBroadcastOperator(t *testing.T) {
	t.Run("Test Chain", func(t *testing.T) {
		bo := NewBroadcastOperator(nil, nil).
			To("room1").
			Except("room2")
		assert.Contains(t, bo.rooms, "room1")
		assert.Contains(t, bo.exceptRooms, "room2")
	})
}
