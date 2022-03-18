package golibemitter

var (
	PacketType = map[string]int{
		"CONNECT":       0,
		"DISCONNECT":    1,
		"EVENT":         2,
		"ACK":           3,
		"CONNECT_ERROR": 4,
		"BINARY_EVENT":  5,
		"BINARY_ACK":    6,
	}
)
