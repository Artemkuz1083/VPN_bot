package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки .env файла")
	}

	host := os.Getenv("DB_HOST")
	port := 5432
	user := "postgres"
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db, nil
}

// CheckUserExists проверяет, существует ли пользователь
func CheckUserExists(userId int) bool {
	db, err := connection()
	if err != nil {
		return false
	}
	defer db.Close()

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE userId=$1)`
	err = db.QueryRow(query, userId).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}
