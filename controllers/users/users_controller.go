package users

import (
	"referral_app/domain/users"
	"referral_app/services"
	"referral_app/utils/errors"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create
// @Tags auth
// @Description register a new user
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body users.RegisterUserRequest true "account info"
// @Success 201 {object} users.User
// @Failure 400 {object} errors.RestErr
// @Failure 401 {object} errors.RestErr
// @Failure 404 {object} errors.RestErr
// @Failure 500 {object} errors.RestErr
// @Failure default {object} errors.RestErr
// @Router /users [post]
func Create(c *gin.Context) {
	var user users.RegisterUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// @Summary Login
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body users.LoginUserRequest true "credentials"
// @Success 200 {string} string "token"
// @Failure 400 {object} errors.RestErr
// @Failure 401 {object} errors.RestErr
// @Failure 404 {object} errors.RestErr
// @Failure 500 {object} errors.RestErr
// @Failure default {object} errors.RestErr
// @Router /users/login [post]
func Login(c *gin.Context) {
	var request users.LoginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	jwtToken, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, jwtToken)
}
