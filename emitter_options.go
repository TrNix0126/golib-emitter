package golibemitter

type EmitterOptions struct {
	Key string
}

func NewEmitterOptions(key string) *EmitterOptions {
	if len(key) == 0 {
		key = "socket.io"
	}
	return &EmitterOptions{Key: key}
}
