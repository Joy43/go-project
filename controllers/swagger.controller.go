package controllers

// Note: Swagger annotations live here to describe the handlers implemented in other files.
// Annotations reference concrete models to produce accurate request/response schemas.

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

// Logout godoc
// @Summary Logout user
// @Description Logout the current user (client-side token removal)
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /logout [post]

// GetProfile godoc
// @Summary Get profile
// @Description Get current user profile
// @Tags Profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]string
// @Router /profile [get]

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

// DeleteAccount godoc
// @Summary Delete account
// @Description Delete user account
// @Tags Profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /profile [delete]

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

// GetAllUsers godoc
// @Summary Get all users
// @Description Admin retrieves all users
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.User
// @Failure 403 {object} map[string]string
// @Router /admin/users [get]

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