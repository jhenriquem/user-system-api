package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jhenriquem/user-system-api/database"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	userID := claims["user_id"]

	response := map[string]any{
		"status": "",
		"userID": userID,
	}

	// Exclusão no banco de dados
	_, err := database.DB.Exec(database.ActionsDB["deleteOne"], userID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response["status"] = "User not found"
	} else {
		w.WriteHeader(http.StatusOK)
		response["status"] = "User deleted"
	}

	// Responsta
	json.NewEncoder(w).Encode(response)
}
