package main

import (
	"time"

	"example.com/quran/db"
	"example.com/quran/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(".env not loaded")
	}

	r := gin.Default()

	db.InitDB()

	configCors := setCORS()
	r.Use(cors.New(configCors))

	routes.RegisterRoutes(r)

	r.Run(":8083")
}

func setCORS() cors.Config {
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://127.0.0.1:5500"}
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Content-Type", "Origin", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	return config
}
