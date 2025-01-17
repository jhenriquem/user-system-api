package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/jhenriquem/user-system-api/database"
	"github.com/jhenriquem/user-system-api/types"

	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	var user types.AuthUser

	// Verificação básica de dados
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and password required", http.StatusBadRequest)
		return
	}

	// Consulta no banco de dados
	var hashPassword string
	var userID int

	err := database.DB.QueryRow(database.ActionsDB["find_by_email"], user.Email).Scan(&userID, &hashPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Validação da senha
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid user or password", http.StatusUnauthorized)
		return
	}

	// Criação do Token
	claims := &Claims{
		UserID: userID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Expira em 24 horas
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Resposta à requisição
	response := map[string]any{
		"status": "Authenticated",
		"token":  token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
