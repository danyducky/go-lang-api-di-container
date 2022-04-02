package controllers

import "go.uber.org/fx"

var Container = fx.Options(
	fx.Provide(NewSignController),
	fx.Provide(NewUserController),
)
