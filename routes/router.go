package routes

import (
	"github.com/gorilla/mux"
	"github.com/jhenriquem/user-system-api/controllers"
	"github.com/jhenriquem/user-system-api/middleware"
)

var Router *mux.Router = mux.NewRouter()

var path map[string]string = map[string]string{
	"register":       "/api/users",
	"authentication": "/api/users/auth",

	"user": "/api/user/",
}

func InitRoutes() {
	Router.HandleFunc(path["authentication"], controllers.AuthUserHandler).Methods("GET")
	Router.HandleFunc(path["register"], controllers.RegisterUserHandler).Methods("POST")

	userRoutes := Router.PathPrefix(path["user"]).Subrouter()
	userRoutes.Use(middleware.TokenAuthentication)

	userRoutes.HandleFunc("/", controllers.GetUserHandler).Methods("GET")
	userRoutes.HandleFunc("/", controllers.UpdateUserHandler).Methods("PUT")
	userRoutes.HandleFunc("/", controllers.DeleteUserHandler).Methods("DELETE")
}
