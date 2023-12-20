-- name: CreateUser :one
INSERT INTO users (
    name, phone_number
) VALUES (
$1, $2
) RETURNING id;

-- name: PhoneNumberExists :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE phone_number = $1
);

-- name: GenerateOTP :exec
UPDATE users 
SET otp = $1, otp_expiration_time = $2 
WHERE phone_number = $3;

-- name: VerifyOTP :one
SELECT otp, otp_expiration_time 
FROM users 
WHERE phone_number = $1;