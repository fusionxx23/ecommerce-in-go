package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fusionxx23/ecommerce-go/http/database"
	"github.com/fusionxx23/ecommerce-go/http/libs"
	"github.com/fusionxx23/ecommerce-go/http/models"
	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

func AuthHandler(s *mux.Router) {
	s.HandleFunc("/refresh", refresh).Methods("GET") // handle refresh token
	s.HandleFunc("/{provider}", signIn)
	s.HandleFunc("/{provider}/callback", callback)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		//redirect to /cart
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func callback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error completing user auth", http.StatusInternalServerError)
		return
	}
	id, err := libs.GenerateRandomID(32)
	if err != nil {
		http.Error(w, "Error generating random ID", http.StatusInternalServerError)
		return
	}
	u := models.User{}
	result := database.DB.Where("email = ?", user.Email).FirstOrCreate(&u, models.User{ // find or create user
		Email:        user.Email,
		RefreshToken: id,
	})
	fmt.Println(result.Error)
	if result.Error != nil {
		fmt.Println(result.Error)
		http.Error(w, "Unable to find or create user", http.StatusInternalServerError)
		return
	}
	if result.Statement.RowsAffected == 1 { // user exist
		u.RefreshToken = id // update refresh token
		if err := database.DB.Save(&u).Error; err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to update user", http.StatusInternalServerError)
			return
		}
	}

	// Create an HTTP cookie with the refresh token
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    id,
		HttpOnly: true,
		Secure:   true, // Set to true in production for HTTPS
		Path:     "/auth/refresh",
		MaxAge:   3600 * 24 * 7, // 7 days
	}
	// Set the cookie in the response
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound)
}
func refresh(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Refreshing JWT Token")
	var user models.User
	result := database.DB.First(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		http.Error(w, "Unable to find user", http.StatusInternalServerError)
		return
	}
	jwtToken, err := libs.CreateJWT(user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to create JWT token", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"access_token": jwtToken,
		"expires_at":   time.Now().Add(86400 * time.Second).Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
