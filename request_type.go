package golibemitter

var RequestType = map[string]int{
	"SOCKETS":           0,
	"ALL_ROOMS":         1,
	"REMOTE_JOIN":       2,
	"REMOTE_LEAVE":      3,
	"REMOTE_DISCONNECT": 4,
	"REMOTE_FETCH":      5,
	"SERVER_SIDE_EMIT":  6,
}
