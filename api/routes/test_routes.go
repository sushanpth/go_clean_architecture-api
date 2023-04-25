package routes

import (
	"clean-architecture-api/api/responses"
	"clean-architecture-api/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestRoutes struct {
	handler infrastructure.Router
}

func NewTestRoutes(handler infrastructure.Router) TestRoutes {
	return TestRoutes{
		handler: handler,
	}
}

func (t TestRoutes) Setup() {
	t.handler.GET("/test", func(c *gin.Context) {
		responses.JSON(c, http.StatusOK, "Hello From Test Route")
	})
}
