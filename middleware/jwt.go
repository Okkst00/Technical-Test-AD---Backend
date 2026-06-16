package middleware

import (
	"backend-api-commerce/repository"
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret-key")

func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(
				w,
				"Authorization token required",
				http.StatusUnauthorized,
			)
			return
		}

		tokenString := strings.Replace(
			authHeader,
			"Bearer ",
			"",
			1,
		)

		// cek blacklist token
		if repository.IsTokenBlacklisted(tokenString) {
			http.Error(
				w,
				"Token already logout",
				http.StatusUnauthorized,
			)
			return
		}

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(
				w,
				"Invalid token",
				http.StatusUnauthorized,
			)
			return
		}

		// ambil claims JWT
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(
				w,
				"Invalid claims",
				http.StatusUnauthorized,
			)
			return
		}

		// simpan data user ke context
		ctx := context.WithValue(
			r.Context(),
			"user_id",
			claims["id"],
		)

		ctx = context.WithValue(
			ctx,
			"user_role",
			claims["role"],
		)

		r = r.WithContext(ctx)

		// lanjut request
		next.ServeHTTP(w, r)
	})
}