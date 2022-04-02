package repositories

import "go.uber.org/fx"

var Container = fx.Options(
	fx.Provide(NewUserRepository),
)
