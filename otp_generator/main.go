// main.go
package main

import (
	"log"

	"github.com/aqaryTest/otp_generator/db"
	"github.com/aqaryTest/otp_generator/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	if err := db.InitDB(); err != nil {
		log.Fatal("Error connectiong to database, check the configuration")
	}

	// Define routes
	r.POST("/api/users", handlers.CreateUserHandler)
	r.POST("/api/users/generateotp", handlers.GenerateOTPHandler)
	r.POST("/api/users/verifyotp", handlers.VerifyOTPHandler)

	// Run the server
	r.Run(":8080")
}
