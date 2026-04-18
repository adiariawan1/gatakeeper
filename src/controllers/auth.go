package controllers

import (
	"net/http"
	"encoding/json"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("110100100");

func LoginHeader(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w,"Method not allow", http.StatusMethodNotAllowed)
		return 
	}

	var creds struct{
		Username string;
		Password string;
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	 
	if err != nil {
		http.Error(w,"format salah",http.StatusBadRequest)
		return
	}

	if creds.Username != "admin" || creds.Password != "admin123" {
		http.Error(w,"password incorrect", http.StatusUnauthorized)
		return
	}

	var claims = jwt.MapClaims{
		"username" : creds.Username,
		"role" : "super admin",
		"exp" : time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)


	if err != nil {
		http.Error(w,"gagal mencetak token ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"token": "` + tokenString + `"}`))
}