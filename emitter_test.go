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
	t.Run("Emit to all with json data", func(t *testing.T) {
		type ExampleJson struct {
			Key     string `json:"key"`
			Message string `json:"message"`
		}
		err = emitter.Emit("hello", &ExampleJson{
			Key:     "this is key",
			Message: "this is message",
		})
		assert.Nil(t, err)
	})
	t.Run("Emit to room", func(t *testing.T) {
		err = emitter.To("room").Emit("hello")
		assert.Nil(t, err)
	})
}
