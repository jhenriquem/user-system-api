package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jhenriquem/user-system-api/database"
	"github.com/jhenriquem/user-system-api/types"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(database.ActionsDB["getAll"])
	if err != nil {
		log.Printf("Database query error: %v", err) // Log do erro
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	users := []types.User{}

	for rows.Next() {

		var u types.User

		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Printf("Row scan error: %v", err) // Log do erro
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err) // Log do erro
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
