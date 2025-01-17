package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jhenriquem/user-system-api/database"
	"github.com/jhenriquem/user-system-api/types"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	userID := claims["user_id"]

	user := types.User{}

	// Recuperação no banco de dados
	err := database.DB.QueryRow(database.ActionsDB["getOne"], userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Printf("Database query error: %v", err) // Log do erro
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Corpo da responsta
	response := map[string]any{
		"status": "Get user data sucessfully",
		"data":   user,
	}

	// Resposta
	json.NewEncoder(w).Encode(response)
}
