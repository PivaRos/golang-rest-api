package structs

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
