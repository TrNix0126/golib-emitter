# Golib Emitter

## Installation

```shell
go get gitlab.com/technixo/golib-emitter
```

Update configuration

```yaml
app.emitter:
  host: localhost
  port: 6379
  database: 0
  username: default
  password: secret
  enableTLS: true
  key: key # default is 'socket.io'
  namespace: ns # default is '/'
```

Update bootstrap fx container

```text
golibemitter.EnableEmitter()
```

## Usage

```go
package emitter

import (
	"gitlab.com/technixo/golib-emitter"
)

type NeedEmitter struct {
}

type ExampleJsonData struct {
	FirstMessage  string `json:"first_message"`
	SecondMessage string `json:"second_message"`
}

// EmitAll emit event to all
func (ne *NeedEmitter) EmitAll() {
	err := golibemitter.Emit("event", "message")
	if err != nil {
		// Error handler
	}
}

// EmitRooms emit event to rooms
func (ne *NeedEmitter) EmitRooms() {
	err := golibemitter.To("room1", "room2").Emit("event", "message")
	if err != nil {
		// Error handler
	}
}

// EmitJSON emit json data message
func (ne *NeedEmitter) EmitJSON() {
	data := ExampleJsonData{
		FirstMessage:  "this is first message",
		SecondMessage: "this is second message",
	}
	err := golibemitter.Emit("event", data)
	if err != nil {
		// Error handler 
	}
}
```
