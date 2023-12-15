package model

import "time"

type User struct {
	ID                int32     `json:"id"`
	Name              string    `json:"name"`
	PhoneNumber       string    `json:"phone_number"`
	OTP               string    `json:"otp"`
	OTPExpirationTime time.Time `json:"otp_expiration_time"`
}
