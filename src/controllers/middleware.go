package controllers

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w,"anda tidak punya kartu akses", http.StatusBadRequest)
			return
		}

		cleanToken := strings.TrimPrefix(token, "Bearer ")

		realToken, err := jwt.Parse(cleanToken, func(token *jwt.Token) (interface{}, error) {
			if _, 	ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return jwtKey, nil
		})
		if err != nil || !realToken.Valid {
            http.Error(w, "token tidak valid atau kedaluwarsa", http.StatusUnauthorized)
            return
        }
		next.ServeHTTP(w, r)
	}
}