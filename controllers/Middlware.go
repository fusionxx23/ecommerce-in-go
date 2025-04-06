package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/libs"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example: Check for a valid token in the Authorization header
		token := r.Header.Get("Authorization")

		var claims jwt.MapClaims
		// Validate the token (you can use your JWT library here)
		parsedToken, err := libs.ParseJWT(token, claims)
		if err != nil {
			fmt.Println("Failed to parse JWT.")
			ctx := r.Context()
			ctx = context.WithValue(ctx, "email", nil) // ensure email is null
			r = r.WithContext(ctx)
			return
		}
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			email := fmt.Sprint(claims["email"])
			ctx := r.Context()
			ctx = context.WithValue(ctx, "email", email)
			r = r.WithContext(ctx)
		} else {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "email", nil) // ensure email is null
			r = r.WithContext(ctx)
		}
		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
