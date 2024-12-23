package golibemitter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
	"strings"
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
		onlyRooms = append(onlyRooms, strings.ToLower(room))
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
		exceptRooms = append(exceptRooms, strings.ToLower(room))
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
		Compress: compress,
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
		Volatile: volatile,
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
	flags := make(map[string]interface{})
	if bo.flags.Compress {
		flags["compress"] = true
	}
	if bo.flags.Volatile {
		flags["volatile"] = true
	}
	pack = append(pack, UID, map[string]interface{}{
		"type": PacketType["EVENT"],
		"data": data,
		"nsp":  bo.broadcastOption.Namespace,
	}, map[string]interface{}{
		"rooms":  bo.rooms,
		"flags":  flags,
		"except": bo.exceptRooms,
	})
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.SetCustomStructTag("json")
	err := enc.Encode(pack)
	if err != nil {
		return fmt.Errorf("broadcast operator: could not encode data: %v", err)
	}
	broadcastChannel := bo.broadcastOption.BroadcastChannel
	if len(bo.rooms) == 1 {
		broadcastChannel = fmt.Sprintf("%s%s#", broadcastChannel, bo.rooms[0])
	}
	return bo.redisClient.Publish(context.Background(), broadcastChannel, buf.Bytes()).Err()
}

func (bo *BroadcastOperator) SocketJoins(rooms ...string) error {
	request := map[string]interface{}{
		"type": PacketType["REMOTE_JOIN"],
		"opts": map[string]interface{}{
			"rooms":  bo.rooms,
			"except": bo.exceptRooms,
		},
		"rooms": bo.roomsToLowerCase(rooms...),
	}
	b, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("socket joins: could not serialize: %v", err)
	}
	return bo.redisClient.Publish(context.Background(), bo.broadcastOption.RequestChannel, b).Err()
}

func (bo *BroadcastOperator) SocketLeave(rooms ...string) error {
	request := map[string]interface{}{
		"type": PacketType["REMOTE_LEAVE"],
		"opts": map[string]interface{}{
			"rooms":  bo.rooms,
			"except": bo.exceptRooms,
		},
		"rooms": bo.roomsToLowerCase(rooms...),
	}
	b, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("socket leave: could not serialize: %v", err)
	}
	return bo.redisClient.Publish(context.Background(), bo.broadcastOption.RequestChannel, b).Err()
}

func (bo *BroadcastOperator) DisconnectSockets(close bool) error {
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
		return fmt.Errorf("disconnect sockets: could not serialize: %v", err)
	}
	return bo.redisClient.Publish(context.Background(), bo.broadcastOption.RequestChannel, b).Err()
}

func (bo BroadcastOperator) roomsToLowerCase(rooms ...string) []string {
	roomsInLowerCase := make([]string, len(rooms))
	for _, room := range rooms {
		roomsInLowerCase = append(roomsInLowerCase, strings.ToLower(room))
	}
	return roomsInLowerCase
}
