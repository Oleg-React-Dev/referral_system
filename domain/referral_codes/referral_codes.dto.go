package referral_codes

import (
	"referral_app/utils/date_utils"
	"referral_app/utils/errors"
	"strings"
)

type ReferralCodeRequest struct {
	ExpiresAt string `json:"expires_at" binding:"required" format:"date-time" example:"2024-11-30 07:45:48.250048+00"`
}

type ReferralCode struct {
	UserId string `json:"user_id,omitempty"`
	Code   string `json:"code"`
	ReferralCodeRequest
}

func (rc *ReferralCode) Validate() *errors.RestErr {
	rc.UserId = strings.TrimSpace(rc.UserId)
	if rc.UserId == "" {
		return errors.NewBadRequestError("invalid user id")
	}

	if err := date_utils.ValidateExpirationDate(rc.ExpiresAt); err != nil {
		return err
	}
	return nil
}
