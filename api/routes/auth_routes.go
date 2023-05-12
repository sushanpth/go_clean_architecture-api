package routes

import (
	"clean-architecture-api/api/controllers"
	"clean-architecture-api/infrastructure"
	"clean-architecture-api/lib"
)

type AuthRoutes struct {
	logger         lib.Logger
	handler        infrastructure.Router
	authController *controllers.AuthController
}

func NewAuthRoutes(
	logger lib.Logger,
	handler infrastructure.Router,
	authController *controllers.AuthController,
) *AuthRoutes {
	return &AuthRoutes{
		logger:         logger,
		handler:        handler,
		authController: authController,
	}
}

func (s *AuthRoutes) Setup() {
	api := s.handler.Group("/api")

	api.POST("/auth/register", s.authController.RegisterUser)
}
