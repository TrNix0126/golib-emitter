package golibemitter

import "fmt"

type BroadcastOptions struct {
	Namespace        string
	BroadcastChannel string
	RequestChannel   string
}

func NewBroadcastOptions(key string, namespace string) *BroadcastOptions {
	if len(key) == 0 {
		key = DefaultKey
	}
	if len(namespace) == 0 {
		namespace = "/"
	}
	broadcastChannel := fmt.Sprintf("%s#%s#", key, namespace)
	requestChannel := fmt.Sprintf("%s-request#%s#", key, namespace)
	return &BroadcastOptions{
		Namespace:        namespace,
		BroadcastChannel: broadcastChannel,
		RequestChannel:   requestChannel,
	}
}
