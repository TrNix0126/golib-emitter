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

import "gitlab.com/technixo/golib-emitter"

type NeedEmitter struct {
	emitter *golibemitter.Emitter
}

// EmitAll emit event to all
func (ne *NeedEmitter) EmitAll()  {
    ne.emitter.Emit("event", "message")
}

// EmitRooms emit event to rooms
func (ne *NeedEmitter) EmitRooms()  {
	ne.emitter.To("room1", "room2").Emit("event", "message")
}
```
