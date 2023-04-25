package routes

import (
	"clean-architecture-api/api/controllers"
	"clean-architecture-api/api/middlewares"
	"clean-architecture-api/infrastructure"
	"clean-architecture-api/lib"
)

type CategoryRoutes struct {
	logger               lib.Logger
	handler              infrastructure.Router
	categoryController   *controllers.CategoryController
	paginationMiddleware middlewares.PaginationMiddleware
}

func NewCategoryRoutes(
	logger lib.Logger,
	handler infrastructure.Router,
	categoryController *controllers.CategoryController,
	paginationMiddleware middlewares.PaginationMiddleware,
) *CategoryRoutes {
	return &CategoryRoutes{
		logger:               logger,
		handler:              handler,
		categoryController:   categoryController,
		paginationMiddleware: paginationMiddleware,
	}
}

func (s *CategoryRoutes) Setup() {
	api := s.handler.Group("/api")

	api.GET("/category", s.paginationMiddleware.Handle(), s.categoryController.GetCategory)
	api.POST("/category", s.categoryController.CreateCategory)
}
