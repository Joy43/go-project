package controllers

import (
	"encoding/json"
	"go-jwt-auth/config"
	"go-jwt-auth/models"
	"go-jwt-auth/utils"
	"net/http"
)

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/yourusername/go-jwt-auth/utils"
// 	"github.com/yourusername/go-jwt-auth/models"
// 	"github.com/yourusername/go-jwt-auth/config"
// )

// Register user
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Hash the password before storing it
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// Save the user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error saving user to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Login user
func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	_ = json.NewDecoder(r.Body).Decode(&loginUser)

	// Retrieve user from the database
	var user models.User
	if err := config.DB.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare the password
	if !utils.ComparePassword(loginUser.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Send the token to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
