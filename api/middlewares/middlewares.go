package middlewares

import "go.uber.org/fx"

// Provides api middlewares container.
var Container = fx.Options(
	fx.Provide(NewJsonMiddleware),
	fx.Provide(NewResponseMiddleware),
	fx.Provide(NewMiddlewares),
)

// Middleware interface.
type Middleware interface {
	Setup()
}

// Application middlewares.
type Middlewares []Middleware

// Creates application middlewares.
// Order matters!
func NewMiddlewares(
	jsonMiddleware JsonMiddleware,
	responseMiddleware ResponseMiddleware,
) Middlewares {
	return Middlewares{
		jsonMiddleware,
		responseMiddleware,
	}
}

// Iterates each middleware and set up it.
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
