package structs

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router      *mux.Router
	MongoClient *mongo.Client
	Env         Env
}

func (a *App) generateTokens(userID string) (accessToken, refreshToken string, err error) {
	// Set expiration times for each token
	accessTokenExpTime := time.Now().Add(15 * time.Minute)    // e.g., 15 minutes for access token
	refreshTokenExpTime := time.Now().Add(7 * 24 * time.Hour) // e.g., 7 days for refresh token

	// Generate Access Token
	accessTokenClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString(a.Env.Jwt_Secret_Key)
	if err != nil {
		return "", "", err
	}

	// Generate Refresh Token
	refreshTokenClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString(a.Env.Jwt_Secret_Key)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a App) refreshToken(oldRefreshToken string) (string, string, error) {

	token, err := jwt.ParseWithClaims(oldRefreshToken, Claims{}, func(token *jwt.Token) (interface{}, error) {
		return a.Env.Jwt_Secret_Key, nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(Claims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("Invalid refresh token")
	}

	// Generate new access and refresh tokens
	newAccessToken, newRefreshToken, err := a.generateTokens(claims.UserID)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
