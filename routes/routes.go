package routes

import (
	"go-jwt-auth/controllers"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetRoutes() *mux.Router {

	r := mux.NewRouter()


	r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/index.html", http.StatusMovedPermanently)
	})

	// ------- Swagger UI -------
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	//  ------- Root route -------
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Go JWT Auth API!"))
	}).Methods("GET")

	// ------ Auth routes ------
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// ------ Profile routes ------
	r.HandleFunc("/profile", controllers.GetProfile).Methods("GET")
	r.HandleFunc("/profile", controllers.UpdateProfile).Methods("PUT")
	r.HandleFunc("/profile", controllers.DeleteAccount).Methods("DELETE")

	// ------ Other routes ------
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")
	r.HandleFunc("/admin/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/comments", controllers.CreateComment).Methods("POST")

	return r
}