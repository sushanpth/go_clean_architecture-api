package routes

import (
	"clean-architecture-api/api/controllers"
	"clean-architecture-api/api/middlewares"
	"clean-architecture-api/infrastructure"
	"clean-architecture-api/lib"
)

type BusinessRoutes struct {
	logger               lib.Logger
	handler              infrastructure.Router
	businessController   *controllers.BusinessController
	paginationMiddleware middlewares.PaginationMiddleware
	authMiddleware       middlewares.FirebaseAuthMiddleware
}

func NewBusinessRoutes(
	logger lib.Logger,
	handler infrastructure.Router,
	businessController *controllers.BusinessController,
	paginationMiddleware middlewares.PaginationMiddleware,
	authMiddleware middlewares.FirebaseAuthMiddleware,
) *BusinessRoutes {
	return &BusinessRoutes{
		logger:               logger,
		handler:              handler,
		businessController:   businessController,
		paginationMiddleware: paginationMiddleware,
		authMiddleware:       authMiddleware,
	}
}

func (s *BusinessRoutes) Setup() {
	api := s.handler.Group("/api").Use(s.authMiddleware.HandleAuthWithRole(""))

	api.GET("/business", s.paginationMiddleware.Handle(), s.businessController.GetBusiness)
	api.POST("/business", s.businessController.CreateBusiness)
}
