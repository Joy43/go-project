package routes

import (
	"go-jwt-auth/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// import (
// 	"go-jwt-auth/controllers"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// SetRoutes sets up the routes for the application
func SetRoutes() *mux.Router {
	r := mux.NewRouter()

	// Root route ("/") handling request to the home page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Go JWT Auth API!"))
	}).Methods("GET")

	// Other routes, such as register, login, etc.
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	return r
}
