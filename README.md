# aqaryTest/otp_generator

An API for registering user and generating otp for them . The database used is postgres

## Prerequisites
A postgres instance running on localhost on port 5432 with the users table.

You can use postgres docker image from Docker Hub with following steps: 
```
docker pull postgres
docker run --name mypg -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```
And Create the following table before running the otp_generator service.
```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    phone_number VARCHAR(20) UNIQUE,
    otp VARCHAR(6),
    otp_expiration_time TIMESTAMP
);
```
Note: 
If you have postgres running with different configuration, change the connection string 
`"postgres://username:password@localhost/database_name?sslmode=disable" `
with necessary changes in InitDB() in db/db.go and sqlc.json. After changing sqlc.json, 
please run `sqlc generate` to generating the queries package automatically.


## How to run 

```
cd aqaryTest/otp_generator
go run main.go
```

## Requests

1. Add user details to DB <br>
**Url** : http://localhost:8080/api/users
**Method** : POST
<br> **body**: 
```json
{
	"name": "Anu",
	"phone_number": "9539909563"
}
```
2. Generate OTP
**Url** : http://localhost:8080/api/users/generateotp
**Method** : POST
<br> **body**: 
```json
{
	"phone_number": "9539909563"
}
```

3. Verify OTP
**Url** : http://localhost:8080/api/users/verifyotp
**Method** : POST
<br> **body**: 
```json
{
	"otp": "1234",
	"phone_number": "9539909563"
}
```
