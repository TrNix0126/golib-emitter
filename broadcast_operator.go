package golibemitter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.com/golibs-starter/golib/log"
)

type BroadcastOperator struct {
	redisClient     *redis.Client
	broadcastOption *BroadcastOptions
	rooms           []string
	exceptRooms     []string
	flags           *BroadcastFlags
}

func NewBroadcastOperator(redisClient *redis.Client, broadcastOption *BroadcastOptions) *BroadcastOperator {
	return &BroadcastOperator{
		redisClient:     redisClient,
		broadcastOption: broadcastOption,
		rooms:           make([]string, 0),
		exceptRooms:     make([]string, 0),
		flags:           new(BroadcastFlags),
	}
}

func (bo *BroadcastOperator) To(rooms ...string) *BroadcastOperator {
	onlyRooms := make([]string, 0)
	for _, room := range rooms {
		onlyRooms = append(onlyRooms, room)
	}
	return &BroadcastOperator{
		redisClient:     bo.redisClient,
		broadcastOption: bo.broadcastOption,
		rooms:           onlyRooms,
		exceptRooms:     bo.exceptRooms,
		flags:           bo.flags,
	}
}

func (bo *BroadcastOperator) In(rooms ...string) *BroadcastOperator {
	return bo.To(rooms...)
}

func (bo *BroadcastOperator) Except(rooms ...string) *BroadcastOperator {
	exceptRooms := make([]string, 0)
	for _, room := range rooms {
		exceptRooms = append(exceptRooms, room)
	}
	return &BroadcastOperator{
		redisClient:     bo.redisClient,
		broadcastOption: bo.broadcastOption,
		rooms:           bo.rooms,
		exceptRooms:     exceptRooms,
		flags:           bo.flags,
	}
}

func (bo *BroadcastOperator) Compress(compress bool) *BroadcastOperator {
	flags := &BroadcastFlags{
		compress: compress,
	}
	return &BroadcastOperator{
		redisClient:     bo.redisClient,
		broadcastOption: bo.broadcastOption,
		rooms:           bo.rooms,
		exceptRooms:     bo.exceptRooms,
		flags:           flags,
	}
}

func (bo *BroadcastOperator) Volatile(volatile bool) *BroadcastOperator {
	flags := &BroadcastFlags{
		volatile: volatile,
	}
	return &BroadcastOperator{
		redisClient:     bo.redisClient,
		broadcastOption: bo.broadcastOption,
		rooms:           bo.rooms,
		exceptRooms:     bo.exceptRooms,
		flags:           flags,
	}
}

func (bo *BroadcastOperator) Emit(event string, args ...interface{}) error {
	data := []interface{}{event}
	data = append(data, args...)
	pack := make([]interface{}, 0)
	pack = append(pack, "emitter")
	pack = append(pack, map[string]interface{}{
		"type": PacketType["EVENT"],
		"data": data,
		"nsp":  bo.broadcastOption.Namespace,
	})
	pack = append(pack, map[string]interface{}{
		"rooms":  bo.rooms,
		"flags":  bo.flags,
		"except": bo.exceptRooms,
	})
	b, err := msgpack.Marshal(pack)
	if err != nil {
		return fmt.Errorf("broadcast operator: could not encode data: %v", err)
	}
	broadcastChannel := bo.broadcastOption.BroadcastChannel
	if len(bo.rooms) == 1 {
		broadcastChannel = fmt.Sprintf("%s%s#", broadcastChannel, bo.rooms[0])
	}
	return bo.redisClient.Publish(context.Background(), broadcastChannel, b).Err()
}

func (bo *BroadcastOperator) SocketJoins(rooms ...string) {
	request := map[string]interface{}{
		"type": PacketType["REMOTE_JOIN"],
		"opts": map[string]interface{}{
			"rooms":  bo.rooms,
			"except": bo.exceptRooms,
		},
		"rooms": rooms,
	}
	b, err := json.Marshal(request)
	if err != nil {
		log.Errorf("socket joins: could not serialize: %v", err)
		return
	}
	bo.redisClient.Publish(context.Background(), bo.broadcastOption.RequestChannel, b)
}

func (bo *BroadcastOperator) SocketLeave(rooms ...string) {
	request := map[string]interface{}{
		"type": PacketType["REMOTE_LEAVE"],
		"opts": map[string]interface{}{
			"rooms":  bo.rooms,
			"except": bo.exceptRooms,
		},
		"rooms": rooms,
	}
	b, err := json.Marshal(request)
	if err != nil {
		log.Errorf("socket joins: could not serialize: %v", err)
		return
	}
	bo.redisClient.Publish(context.Background(), bo.broadcastOption.RequestChannel, b)
}

func (bo BroadcastOperator) DisconnectSockets(close bool) {
	request := map[string]interface{}{
		"type": PacketType["REMOTE_DISCONNECT"],
		"opts": map[string]interface{}{
			"rooms":  bo.rooms,
			"except": bo.exceptRooms,
		},
		"close": close,
	}
	b, err := json.Marshal(request)
	if err != nil {
		log.Errorf("socket joins: could not serialize: %v", err)
		return
	}
	bo.redisClient.Publish(context.Background(), bo.broadcastOption.RequestChannel, b)
}
