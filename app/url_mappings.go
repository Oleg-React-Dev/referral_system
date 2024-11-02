package app

import (
	"referral_app/controllers/auth"
	"referral_app/controllers/ping"
	"referral_app/controllers/referral_code"
	"referral_app/controllers/users"

	_ "referral_app/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func mapUrls() {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.POST("/users/login", users.Login)

	router.POST("/referral-code", auth.TokenVerifyMiddleWare(referral_code.Create))
	router.DELETE("/referral-code", auth.TokenVerifyMiddleWare(referral_code.Delete))
	router.GET("/referral-code/:email", referral_code.GetReferralCodeByEmail)

	router.GET("/referrals", auth.TokenVerifyMiddleWare(referral_code.GetReferralsByReferrer))
}
