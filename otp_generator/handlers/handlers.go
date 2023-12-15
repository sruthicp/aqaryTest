package handlers

import (
	"database/sql"
	"fmt"

	"github.com/aqaryTest/otp_generator/db"
	"github.com/aqaryTest/otp_generator/model"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler handles the creation of a new user
func CreateUserHandler(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if the phone number already exists
	exists, err := db.PhoneNumberExists(c, user.PhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error checking phone number existence: %s", err.Error())})
		return
	}
	if exists {
		c.JSON(400, gin.H{"error": "Phone number already exists"})
		return
	}

	// Create the user in the database
	err = db.CreateUser(c, &user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(200, user)
}

// GenerateOTPHandler generates OTP for a user
func GenerateOTPHandler(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if the phone number exists
	exists, err := db.PhoneNumberExists(c, request.PhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error checking phone number existence: %s", err.Error())})
		return
	}
	if !exists {
		c.JSON(404, gin.H{"error": "Phone number not registered !!"})
		return
	}

	// Generate and save OTP for the user
	otp, err := db.GenerateOTP(c, request.PhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating OTP"})
		return
	}

	c.JSON(200, gin.H{"message": "OTP generated successfully", "otp": otp})
}

// VerifyOTPHandler verifies OTP for a user
func VerifyOTPHandler(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone_number"`
		OTP         string `json:"otp"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Verify OTP for the user
	err := db.VerifyOTP(c, request.PhoneNumber, request.OTP)
	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "Phone number not registered!!"})
		return
	} else if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Error verifying OTP: %s", err.Error())})
		return
	}

	c.JSON(200, gin.H{"message": "OTP verified successfully"})
}
