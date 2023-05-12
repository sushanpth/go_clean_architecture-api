package controllers

import (
	"clean-architecture-api/api/responses"
	"clean-architecture-api/lib"
	"clean-architecture-api/models"
	"clean-architecture-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.UserService
	logger  lib.Logger
}

func NewAuthController(
	service *services.UserService,
	logger lib.Logger,
) *AuthController {
	return &AuthController{
		service: service,
		logger:  logger,
	}
}

func (ac *AuthController) RegisterUser(c *gin.Context) {

	// get data from body
	type UserData struct {
		Name     string `json:"name" form:"name"`
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"email"`
	}

	var data UserData
	err := c.ShouldBindJSON(&data)

	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Invalid request")
		return
	}
	// create hash password
	hash, err := ac.service.HashPassword(data.Password)

	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to generate password hash.")
		return
	}

	// save user to db
	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hash,
	}
	err = ac.service.Create(&user)

	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Invalid request")
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "OK")

}

// Login logs in a user
func (ac *AuthController) Login(c *gin.Context) {
	// get info from body
	var data struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"email"`
	}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to read user data")
		return
	}

	// get user from db
	user, err := ac.service.GetUserByEmail(data.Email)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Invalid email or password! find user")
		return
	}
	ac.logger.Info(user.Password, " USER ", data.Password)
	// validate password
	isValid := ac.service.CompareHashAndPassword(user.Password, data.Password)
	if !isValid {
		responses.ErrorJSON(c, http.StatusBadRequest, "Invalid email or password!")
		return
	}

	// generate JWT Token
	token, err := ac.service.GenerateJWTToken(user)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Invalid email or password! token err")
		return
	}

	// send response
	responses.SuccessJSON(c, http.StatusOK, gin.H{"success": "OK", "token": token})

}
