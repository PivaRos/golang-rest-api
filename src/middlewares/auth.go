package middlewares

import (
	"my-rest-api/src/structs"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &structs.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return structs.Jwt_Secret_KeyExported, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if _, ok := token.Claims.(*structs.Claims); ok && token.Valid {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
