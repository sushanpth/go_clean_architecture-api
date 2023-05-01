package controllers

import (
	"clean-architecture-api/api/middlewares"
	"clean-architecture-api/api/responses"
	"clean-architecture-api/lib"
	"clean-architecture-api/services"
	"clean-architecture-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BusinessController struct {
	service         *services.BusinessService
	logger          lib.Logger
	dbTrxMiddleware middlewares.DBTransactionMiddleware
}

func NewBusinessController(
	service *services.BusinessService,
	logger lib.Logger,
	dbTrxMiddleware middlewares.DBTransactionMiddleware,
) *BusinessController {
	return &BusinessController{
		service:         service,
		logger:          logger,
		dbTrxMiddleware: dbTrxMiddleware,
	}
}

type body struct {
	Name       string           `json:"name" form:"name"`
	Location   string           `json:"location" form:"location"`
	Categories []lib.BinaryUUID `json:"categories"`
}

func (bc *BusinessController) GetBusiness(c *gin.Context) {
	businesses, err := bc.service.SetPaginationScope(utils.Paginate(c)).GetAll()

	if err != nil {
		bc.logger.Error(err)
	}
	responses.JSONWithPagination(c, http.StatusOK, businesses)
}

func (bc *BusinessController) CreateBusiness(c *gin.Context) {
	var data body
	if err := c.ShouldBindJSON(&data); err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed get request data!")
		return
	}
	bc.logger.Info(data)
	responses.SuccessJSON(c, http.StatusOK, "OK")
}
