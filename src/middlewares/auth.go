package middlewares

import (
	"my-rest-api/src/structs"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthenticateMiddleware(roles []structs.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
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

			if claim, ok := token.Claims.(*structs.Claims); ok && token.Valid {
				var found bool = false
				for _, value := range roles {
					if value == claim.Role {
						found = true
					}
				}
				if found {
					next.ServeHTTP(w, r)
					return
				}
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		})
	}
}
