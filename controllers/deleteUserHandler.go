package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhenriquem/user-system-api/database"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var response map[string]any

	_, err := database.DB.Exec(database.ActionsDB["deleteOne"], params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response = map[string]any{
			"status": "User not found",
			"userID": params["id"],
		}
	} else {
		w.WriteHeader(http.StatusOK)
		response = map[string]any{
			"status": "User deleted",
			"userID": params["id"],
		}
	}

	json.NewEncoder(w).Encode(response)
}
