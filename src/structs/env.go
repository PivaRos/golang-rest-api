package structs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	MONGO_URI                string
	Jwt_Secret_Key           string
	Access_Token_Expiration  time.Duration
	Refresh_Token_Expiration time.Duration
}

func (e *Env) InitEnv() {
	EnvErr := godotenv.Load("../.env")
	if EnvErr != nil {
		log.Fatalf("Error loading .env file: %v", EnvErr)
	}
	e.MONGO_URI = os.Getenv("MONGO_URI")
	if e.MONGO_URI == "" {
		log.Fatalln("Error: no uri was found in env file")
	}
	e.Jwt_Secret_Key = os.Getenv("JWT_SECRET_KEY")
	if e.Jwt_Secret_Key == "" {
		log.Fatalln("Error: no JWT_SECRET_KEY was found in env file")
	}
	Access_Token_Expiration, AccessExpirationErr := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRATION"))
	if AccessExpirationErr != nil || Access_Token_Expiration == time.Duration(0) {
		log.Fatalln("Error: no ACCESS_TOKEN_EXPIRATION was found in env file")
	}
	e.Access_Token_Expiration = Access_Token_Expiration
	Refresh_Token_Expiration, RefreshExpirationErr := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRATION"))
	if RefreshExpirationErr != nil || Refresh_Token_Expiration == time.Duration(0) {
		log.Fatalln("Error: no REFRESH_TOKEN_EXPIRATION was found in env file")
	}
	e.Refresh_Token_Expiration = Refresh_Token_Expiration
}
