package app

import (
	"os"
	"referral_app/config"
	"referral_app/datasources/postegresql/users_db"
	"referral_app/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	if err := config.LoadEnv(); err != nil {
		logger.Error("error occurred while loading env file:", err)
		panic(err)
	}

	if err := users_db.InitDB(); err != nil {
		logger.Error("error occurred while initializing database:", err)
		panic(err)
	}

	mapUrls()

	logger.Info("about to start the application...")

	if err := router.Run(os.Getenv("PORT")); err != nil {
		logger.Error("error occurred while running http server:", err)
		panic(err)
	}

}
