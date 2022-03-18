package golibemitter

import (
	"gitlab.com/golibs-starter/golib/config"
)

// EmitterProperties represents websocket emitter properties
type EmitterProperties struct {
	Key       string
	Namespace string
	Host      string
	Port      int
	Username  string
	Password  string
	Database  int
	EnableTLS bool
}

// NewEmitterProperties return a new EmitterProperties instance
func NewEmitterProperties(loader config.Loader) (*EmitterProperties, error) {
	props := EmitterProperties{}
	err := loader.Bind(&props)
	return &props, err
}

// Prefix return config prefix
func (t *EmitterProperties) Prefix() string {
	return "app.emitter"
}
