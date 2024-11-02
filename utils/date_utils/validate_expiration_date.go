package date_utils

import (
	"referral_app/utils/errors"
	"strings"
	"time"
)

var apiDbLayout = "2006-01-02 15:04:05.999999+00"

func ValidateExpirationDate(expiresAt string) *errors.RestErr {
	expiresAt = strings.TrimSpace(expiresAt)

	if expiresAt == "" {
		return errors.NewBadRequestError("invalid expiration date: cannot be empty")
	}

	parsedTime, err := time.Parse(apiDbLayout, expiresAt)
	if err != nil {
		return errors.NewBadRequestError("invalid expiration date: must be in 'YYYY-MM-DD HH:MM:SS.mmmmmm+00' format")
	}

	if !parsedTime.After(time.Now()) {
		return errors.NewBadRequestError("invalid expiration date: must be in the future")
	}
	return nil
}
