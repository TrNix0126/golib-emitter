package golibemitter

const DefaultKey = "socket.io"

type EmitterOptions struct {
	Key string
}

func NewEmitterOptions(key string) *EmitterOptions {
	if len(key) == 0 {
		key = DefaultKey
	}
	return &EmitterOptions{Key: key}
}
