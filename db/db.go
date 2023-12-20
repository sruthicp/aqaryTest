package db

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/aqaryTest/otp_generator/model"
	"github.com/aqaryTest/otp_generator/queries"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func InitDB() error {
	var err error
	// connection string to postgres
	connString := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return err
	}

	err = db.Ping(context.Background())
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return err
	}

	fmt.Println("Connected to the database", db)
	return nil

}

// CreateUser inserts a new user into the database
func CreateUser(ctx *gin.Context, user *model.User) error {
	Q := queries.New(db)
	args := queries.CreateUserParams{
		Column1: sql.NullString{String: user.Name, Valid: true},
		Column2: sql.NullString{String: user.PhoneNumber, Valid: true},
	}
	id, err := Q.CreateUser(ctx, args)
	user.ID = id

	return err
}

// PhoneNumberExists checks if a phone number already exists in the database
func PhoneNumberExists(ctx *gin.Context, phoneNumber string) (bool, error) {
	Q := queries.New(db)
	arg := sql.NullString{String: phoneNumber, Valid: true}
	exists, err := Q.PhoneNumberExists(ctx, arg)

	return exists, err
}

// GenerateOTP generates a new OTP for the user and sets its expiration time
func GenerateOTP(ctx *gin.Context, phoneNumber string) (string, error) {
	// generating a random 4 digit otp
	otp := fmt.Sprint(rand.Intn(9000) + 1000)

	expirationTime := time.Now().Add(time.Minute)
	Q := queries.New(db)
	args := queries.GenerateOTPParams{
		Column1: sql.NullString{String: otp, Valid: true},
		Column2: sql.NullTime{Time: expirationTime, Valid: true},
		Column3: sql.NullString{String: phoneNumber, Valid: true},
	}
	err := Q.GenerateOTP(ctx, args)
	return otp, err
}

// VerifyOTP verifies the OTP for a user
func VerifyOTP(ctx *gin.Context, phoneNumber, otp string) error {
	var storedOTP string
	var expirationTime time.Time

	Q := queries.New(db)
	arg := sql.NullString{String: phoneNumber, Valid: true}
	row, err := Q.VerifyOTP(ctx, arg)
	storedOTP = row.Otp.String
	expirationTime = row.OtpExpirationTime.Time

	if err == sql.ErrNoRows {
		return sql.ErrNoRows // Phone number not found
	} else if err != nil {
		return err
	}

	if otp != storedOTP {
		return fmt.Errorf("OTP not matching !")
	}

	expirationTimeString := expirationTime.Format("2006-01-02 15:04:05")
	currentTimeString := time.Now().Format("2006-01-02 15:04:05")

	if expirationTimeString < currentTimeString {
		return fmt.Errorf("OTP expired !")
	}

	return nil
}
