package golibemitter

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Emitter struct {
	emitterOption    *EmitterOptions
	broadcastOptions *BroadcastOptions
	redisClient      *redis.Client
}

func NewEmitterFromProperties(properties *EmitterProperties) (*Emitter, error) {
	config := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", properties.Host, properties.Port),
		Username: properties.Username,
		Password: properties.Password,
		DB:       properties.Database,
	}
	if properties.EnableTLS {
		config.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}
	redisClient := redis.NewClient(config)
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		return nil, err
	}
	emitterOptions := NewEmitterOptions(properties.Key)
	broadcastOptions := NewBroadcastOptions(emitterOptions.Key, properties.Namespace)
	return &Emitter{
		emitterOption:    emitterOptions,
		broadcastOptions: broadcastOptions,
		redisClient:      redisClient,
	}, nil
}

func NewEmitter(emitterOption *EmitterOptions, broadcastOptions *BroadcastOptions, redisClient *redis.Client) *Emitter {
	return &Emitter{
		emitterOption:    emitterOption,
		broadcastOptions: broadcastOptions,
		redisClient:      redisClient,
	}
}

func (e *Emitter) Of(namespace string) *Emitter {
	broadcastOptions := &BroadcastOptions{
		Namespace:        namespace,
		BroadcastChannel: e.broadcastOptions.BroadcastChannel,
		RequestChannel:   e.broadcastOptions.RequestChannel,
	}
	return NewEmitter(e.emitterOption, broadcastOptions, e.redisClient)
}

func (e *Emitter) Emit(event string, args ...interface{}) error {
	return NewBroadcastOperator(e.redisClient, e.broadcastOptions).Emit(event, args...)
}

func (e *Emitter) To(rooms ...string) *BroadcastOperator {
	return NewBroadcastOperator(e.redisClient, e.broadcastOptions).To(rooms...)
}

func (e *Emitter) In(rooms ...string) *BroadcastOperator {
	return e.To(rooms...)
}

func (e *Emitter) Except(rooms ...string) *BroadcastOperator {
	return NewBroadcastOperator(e.redisClient, e.broadcastOptions).Except(rooms...)
}
