package referral_code

import (
	"net/http"
	"referral_app/domain/referral_codes"
	"referral_app/services"
	"referral_app/utils/errors"
	"regexp"

	"github.com/gin-gonic/gin"
)

// @Summary Create referral code
// @Security ApiKeyAuth
// @Tags referral code
// @Description Create referral code with expiration date
// @ID create-code
// @Accept json
// @Produce json
// @Param referralCode body referral_codes.ReferralCodeRequest true "Referral code expiration details"
// @Success 201 {object} referral_codes.ReferralCode
// @Failure 400,404 {object} errors.RestErr
// @Failure 500 {object} errors.RestErr
// @Failure default {object} errors.RestErr
// @Router /referral-code [post]
func Create(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		restErr := errors.NewUnauthorizedError("user ID not found in context")
		c.JSON(restErr.Status, restErr)
		return
	}

	var rc referral_codes.ReferralCode
	rc.UserId = userId.(string)
	if err := c.ShouldBindJSON(&rc); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.ReferralService.CreateReferralCode(rc)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// @Summary Delete referral code
// @Security ApiKeyAuth
// @Tags referral code
// @Description delete referral code
// @ID delete-code
// @Accept  json
// @Produce  json
// @Success 200 {string} {"status": "deleted"}
// @Failure 400,404 {object} errors.RestErr
// @Failure 500 {object} errors.RestErr
// @Failure default {object} errors.RestErr
// @Router /referral-code [delete]
func Delete(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		restErr := errors.NewUnauthorizedError("user ID not found in context")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := services.ReferralService.DeleteReferralCode(userId.(string)); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// @Summary Get referral code by email
// @Tags referral code
// @Description get referral code by email
// @ID get-code
// @Accept  json
// @Produce  json
// @Param email path string true "Email Address"
// @Success 200 {object} referral_codes.ReferralCode
// @Failure 400,404 {object} errors.RestErr
// @Failure 500 {object} errors.RestErr
// @Failure default {object} errors.RestErr
// @Router /referral-code/{email} [get]
func GetReferralCodeByEmail(c *gin.Context) {
	email := c.Param("email")
	if !isEmailValid(email) {
		restErr := errors.NewBadRequestError("invalid email address")
		c.JSON(restErr.Status, restErr)
		return
	}

	referralCode, err := services.ReferralService.GetReferralCodeByEmail(email)
	if err != nil {
		restErr := errors.NewNotFoundError("user not found")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, referralCode)
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

// @Summary Get referrals by referrer
// @Security ApiKeyAuth
// @Tags referral code
// @Description get all referrals belong to referrer
// @ID all-referrals
// @Accept  json
// @Produce  json
// @Success 200 {object} users.Users
// @Failure 400,404 {object} errors.RestErr
// @Failure 500 {object} errors.RestErr
// @Failure default {object} errors.RestErr
// @Router /referrals [get]
func GetReferralsByReferrer(c *gin.Context) {
	referrerId, exists := c.Get("user_id")
	if !exists {
		restErr := errors.NewUnauthorizedError("user ID not found in context")
		c.JSON(restErr.Status, restErr)
		return
	}

	referrals, err := services.ReferralService.GetReferralsByReferrerId(referrerId.(string))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, referrals)
}
