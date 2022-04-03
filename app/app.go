package app

import "go.uber.org/fx"

// Provides application container.
var Container = fx.Options(
	fx.Provide(NewDatabase),
	fx.Provide(NewConfig),
	fx.Provide(NewRequestHandler),
	fx.Provide(NewMapper),
)
