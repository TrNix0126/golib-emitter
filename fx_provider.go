package golibemitter

import (
	"gitlab.com/golibs-starter/golib"
	"go.uber.org/fx"
)

func EnableEmitter() fx.Option {
	return fx.Options(
		golib.ProvideProps(NewEmitterProperties),
		fx.Provide(NewEmitterFromProperties),
	)
}
