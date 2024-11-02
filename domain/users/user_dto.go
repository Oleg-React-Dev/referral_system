package users

import (
	"referral_app/utils/errors"
	"strings"
)

type User struct {
	UserId       string `json:"user_id" example:"c4f32ef1-8b8e-48c6-95c2-0fe8fbeefddd"`
	Email        string `json:"email" example:"user@example.com"`
	Password     string `json:"-" binding:"required" db:"password"`
	CreatedAt    string `json:"created_at" example:"2024-10-31 07:45:48.250048+00"`
	UpdatedAt    string `json:"updated_at" example:"2024-10-31 07:45:48.250048+00"`
	ReferralCode string `json:"referral_code,omitempty" example:"ABC123"`
}

type Users []User

func (u *User) Validate() *errors.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
