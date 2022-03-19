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
```

Update bootstrap fx container

```text
golibemitter.EnableEmitter()
```

## Usage

```go
package emitter

import (
	"github.com/go-playground/locales/da"
	"gitlab.com/technixo/golib-emitter"
)

type NeedEmitter struct {
	emitter *golibemitter.Emitter
}

type ExampleJsonData struct {
	FirstMessage  string `json:"first_message"`
	SecondMessage string `json:"second_message"`
}

// EmitAll emit event to all
func (ne *NeedEmitter) EmitAll() {
	ne.emitter.Emit("event", "message")
}

// EmitRooms emit event to rooms
func (ne *NeedEmitter) EmitRooms() {
	ne.emitter.To("room1", "room2").Emit("event", "message")
}

// EmitJSON emit json data message
func (ne *NeedEmitter) EmitJSON() {
	data := ExampleJsonData{
		FirstMessage:  "this is first message",
		SecondMessage: "this is second message",
	}
	ne.emitter.Emit("event", data)
}
```
