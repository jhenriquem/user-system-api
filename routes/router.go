package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhenriquem/user-system-api/controllers"
)

var Router *mux.Router = mux.NewRouter()

var path map[string]string = map[string]string{
	"register":       "/users",
	"authentication": "/users/auth",

	"getUser":        "/users/{id}",
	"updateUserData": "/users/{id}",
	"deleteUser":     "/users/{id}",
}

func InitRoutes() {
	Router.HandleFunc(path["authentication"], controllers.AuthUserHandler).Methods("GET")

	Router.HandleFunc(path["getUser"], controllers.GetUserHandler).Methods("GET")

	Router.HandleFunc(path["register"], controllers.RegisterUserHandler).Methods("POST")

	Router.HandleFunc(path["updateUserData"], func(w http.ResponseWriter, r *http.Request) {}).Methods("PUT")

	Router.HandleFunc(path["deleteUser"], controllers.DeleteUserHandler).Methods("DELETE")
}
