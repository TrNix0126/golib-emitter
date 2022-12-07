package golibemitter

import (
	"gitlab.com/golibs-starter/golib"
	"go.uber.org/fx"
)

var emitter *Emitter

func EnableEmitter() fx.Option {
	return fx.Options(
		golib.ProvideProps(NewEmitterProperties),
		fx.Provide(NewEmitterFromProperties),
		fx.Invoke(func(e *Emitter) {
			emitter = e
		}),
	)
}

func Emit(event string, args ...interface{}) error {
	return NewBroadcastOperator(emitter.redisClient, emitter.broadcastOptions).Emit(event, args...)
}

func To(rooms ...string) *BroadcastOperator {
	return NewBroadcastOperator(emitter.redisClient, emitter.broadcastOptions).To(rooms...)
}
