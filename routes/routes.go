package routes

import (
	"net/http"

	"../api"
	"../utils/auth"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/api/register", api.CreateUser).Methods("POST")
	r.HandleFunc("/api/login", api.LoginUser).Methods("POST")

	// Auth route
	s := r.PathPrefix("/api/auth").Subrouter()
	s.Use(auth.JwtVerify)

	// Users
	s.HandleFunc("/user", api.FetchUsers).Methods("GET")
	s.HandleFunc("/user/{id}", api.GetUser).Methods("GET")
	s.HandleFunc("/user/{id}", api.UpdateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", api.DeleteUser).Methods("DELETE")

	// Exercises
	s.HandleFunc("/exercise", api.FetchExercises).Methods("GET")
	s.HandleFunc("/exercise", api.CreateExercise).Methods("POST")
	s.HandleFunc("/user/{id}/exercise", api.FetchUserExercises).Methods("GET")
	s.HandleFunc("/user/{id}/dates", api.FetchUserDates).Methods("GET")
	s.HandleFunc("/user/{id}/exercise/{date}", api.FetchUserDateExercises).Methods("GET")
	s.HandleFunc("/exercise/{id}", api.GetExercise).Methods("GET")
	s.HandleFunc("/exercise/{id}", api.UpdateExercise).Methods("PUT")
	s.HandleFunc("/exercise/{id}", api.DeleteExercise).Methods("DELETE")
	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
