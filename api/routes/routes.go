package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(NewTestRoutes),
	fx.Provide(NewCategoryRoutes),
	fx.Provide(NewBusinessRoutes),
	fx.Provide(NewAuthRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes *UserRoutes,
	testRoutes TestRoutes,
	categoryRoutes *CategoryRoutes,
	businessRoutes *BusinessRoutes,
	authRoutes *AuthRoutes,
) Routes {
	return Routes{
		userRoutes,
		testRoutes,
		categoryRoutes,
		businessRoutes,
		authRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
