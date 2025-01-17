package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhenriquem/user-system-api/database"
	"github.com/jhenriquem/user-system-api/types"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user types.User
	json.NewDecoder(r.Body).Decode(&user)

	IsValid := true

	// Verifica se todos os valores foram passados
	if user.Name == "" || user.Password == "" || user.Email == "" {
		IsValid = false
	}

	if !IsValid {
		http.Error(w, "Incomplete or missing Data", http.StatusBadRequest)
		return
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Verifica se o usuário já existe
	_, err = database.DB.Query(database.ActionsDB["find_by_email"], user.Email)
	if err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Criação no banco de dados
	err = database.DB.QueryRow(database.ActionsDB["create"], user.Name, user.Email, hashedPassword).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Create user error", http.StatusInternalServerError)
		return
	}

	// Resposta bem sucedia
	response := map[string]string{
		"status": "User created successfully",
		"userId": fmt.Sprint(user.ID),
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
