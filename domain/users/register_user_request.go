package users

type RegisterUserRequest struct {
	Email        string `json:"email" binding:"required" example:"user@example.com"`
	Password     string `json:"password" binding:"required"`
	ReferralCode string `json:"referral_code,omitempty"`
}
