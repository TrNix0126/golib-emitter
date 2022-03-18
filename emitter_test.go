package golibemitter

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmitter_Emit(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	assert.Nil(t, redisClient.Ping(context.Background()).Err())
	emitter, err := NewEmitterFromProperties(&EmitterProperties{
		Host: "localhost",
		Port: 6379,
	})
	assert.Nil(t, err)
	t.Run("Emit to all", func(t *testing.T) {
		err = emitter.Emit("hello", "world")
		assert.Nil(t, err)
	})
	t.Run("Emit to room", func(t *testing.T) {
		err = emitter.To("room").Emit("hello")
		assert.Nil(t, err)
	})
}
