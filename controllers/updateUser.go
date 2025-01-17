package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jhenriquem/user-system-api/database"
	"github.com/jhenriquem/user-system-api/types"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user types.User
	json.NewDecoder(r.Body).Decode(&user)

	claims := r.Context().Value("claims").(jwt.MapClaims)
	userID := claims["user_id"]

	// Hash da nova senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec(database.ActionsDB["updateOne"], user.Name, user.Email, hashedPassword, userID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Update user error", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"status": "Update user sucessfully",
		"userId": userID,
	}

	json.NewEncoder(w).Encode(response)
}
