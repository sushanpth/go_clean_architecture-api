package bootstrap

import (
	"clean-architecture-api/api/controllers"
	"clean-architecture-api/api/middlewares"
	"clean-architecture-api/api/routes"
	"clean-architecture-api/infrastructure"
	"clean-architecture-api/lib"
	"clean-architecture-api/repository"
	"clean-architecture-api/seeds"
	"clean-architecture-api/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	services.Module,
	repository.Module,
	infrastructure.Module,
	middlewares.Module,
	lib.Module,
	seeds.Module,
)
