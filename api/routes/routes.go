package routes

import "go.uber.org/fx"

// Provides application route groups container.
var Container = fx.Options(
	fx.Provide(NewSignRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewRoutes),
)

// Route interface.
type Route interface {
	Setup()
}

// Application routes.
type Routes []Route

// Creates application route groups.
func NewRoutes(
	signRoutes SignRoutes,
	userRoutes UserRoutes,
) Routes {
	return Routes{
		signRoutes,
		userRoutes,
	}
}

// Iterates each route group and set up it.
func (r Routes) Setup() {
	for _, router := range r {
		router.Setup()
	}
}
