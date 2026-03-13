package controllers

import (
	"encoding/json"
	"go-jwt-auth/config"
	"go-jwt-auth/models"
	"go-jwt-auth/utils"
	"net/http"
)

// Register godoc
// @Summary Register user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "Register data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /register [post]
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "Login data (email + password)"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	_ = json.NewDecoder(r.Body).Decode(&loginUser)

	// ----- Retrieve user from the database ---------
	var user models.User
	if err := config.DB.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// --------- Compare the password ---------
	if !utils.ComparePassword(loginUser.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// ------Generate JWT token ---------
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// --------- Send the token to the client ---------
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Logout godoc
// @Summary Logout user
// @Description Logout the current user (client-side token removal)
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
	// Since JWT is stateless, we can't invalidate the token on the server side.
	// The client should simply delete the token on logout.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}	

// ----- Refresh JWT token ---------
func RefreshToken(w http.ResponseWriter, r *http.Request) {	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Token refreshed successfully"})
}

// GetProfile godoc
// @Summary Get profile
// @Description Get current user profile
// @Tags Profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]string
// @Router /profile [get]
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// Get the username from the JWT claims (set by middleware)
	username := r.Context().Value("username").(string)

	// Retrieve user information from the database
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// UpdateProfile godoc
// @Summary Update profile
// @Description Update user profile
// @Tags Profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body models.User true "Profile data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /profile [put]
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// Get the username from the JWT claims (set by middleware)
	username := r.Context().Value("username").(string)

	// Retrieve user information from the database
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the updated user information from the request body
	var updatedUser models.User
	_ = json.NewDecoder(r.Body).Decode(&updatedUser)

	// Update user fields (except password and role)
	user.Email = updatedUser.Email

	// Save the updated user to the database
	if err := config.DB.Save(&user).Error; err != nil {
		http.Error(w, "Error updating user profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// ChangePassword godoc
// @Summary Change password
// @Description Change current user's password
// @Tags Profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body map[string]string true "{new_password: string}"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /profile/password [put]
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Get the username from the JWT claims (set by middleware)
	username := r.Context().Value("username").(string)

	// Retrieve user information from the database
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the new password from the request body
	var passwordData struct {
		NewPassword string `json:"new_password"`
	}
	_ = json.NewDecoder(r.Body).Decode(&passwordData)

	// Hash the new password before storing it
	hashedPassword, err := utils.HashPassword(passwordData.NewPassword)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// Save the updated user to the database
	if err := config.DB.Save(&user).Error; err != nil {
		http.Error(w, "Error changing password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password changed successfully"})
}

// DeleteAccount godoc
// @Summary Delete account
// @Description Delete user account
// @Tags Profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /profile [delete]
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// Get the username from the JWT claims (set by middleware)
	username := r.Context().Value("username").(string)

	// Retrieve user information from the database
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Delete the user from the database
	if err := config.DB.Delete(&user).Error; err != nil {
		http.Error(w, "Error deleting user account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Account deleted successfully"})
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Admin retrieves all users
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.User
// @Failure 403 {object} map[string]string
// @Router /admin/users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Get the username from the JWT claims (set by middleware)
	username := r.Context().Value("username").(string)

	// Retrieve user information from the database
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if the user has admin role
	if user.Role != "admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Retrieve all users from the database
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// CreatePost godoc
// @Summary Create post
// @Description Create a new post
// @Tags Posts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param post body models.Post true "Post data"
// @Success 201 {object} models.Post
// @Failure 400 {object} map[string]string
// @Router /posts [post]
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Get the username from the JWT claims (set by middleware)
	username := r.Context().Value("username").(string)

	// Retrieve user information from the database
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the new post from the request body
	var post models.Post
	_ = json.NewDecoder(r.Body).Decode(&post)

	// Set the post's user ID
	post.UserID = user.ID

	// Save the new post to the database
	if err := config.DB.Create(&post).Error; err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

// CreateComment godoc
// @Summary Create comment
// @Description Add a comment to a post
// @Tags Comments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment data"
// @Success 201 {object} models.Comment
// @Failure 400 {object} map[string]string
// @Router /comments [post]
func CreateComment(w http.ResponseWriter, r *http.Request) {
	// Get the username from the JWT claims (set by middleware)
	username := r.Context().Value("username").(string)

	// Retrieve user information from the database
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the new comment from the request body
	var comment models.Comment
	_ = json.NewDecoder(r.Body).Decode(&comment)

	// Set the comment's user ID
	comment.UserID = user.ID

	// Save the new comment to the database
	if err := config.DB.Create(&comment).Error; err != nil {
		http.Error(w, "Error creating comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the home page!"})
}