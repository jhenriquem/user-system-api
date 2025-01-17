package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB, err   = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	ActionsDB = map[string]string{
		"init":          "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT , password TEXT)",
		"create":        "INSERT INTO users (name, email, password) VALUES ($1 ,$2 ,$3) RETURNING id",
		"find_by_email": "SELECT id, password FROM users WHERE email = $1",
		"getOne":        "SELECT id,name, email FROM users WHERE id = $1",
		"updateOne":     "UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4",
		"deleteOne":     "DELETE FROM users WHERE id = $1",
	}
)

func InitDatabase() {
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Database online")

	// Criar a tabela caso não exista
	_, err = DB.Exec(ActionsDB["init"])
	if err != nil {
		log.Fatal(err)
	}
}
