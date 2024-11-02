package main

import app "referral_app/app"

// @title Simple API Example
// @version 1.0
// @description Example API for testing Swagger generation

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.StartApp()
}
