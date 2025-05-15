package database

//Здесь реализовано подключение к БД, вообще впринципе что связано с БД тут

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Осуществляет подключение к БД
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
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE telegram_id=$1)`
	err = db.QueryRow(query, userId).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}

// Добавление нового пользователя в БД
func AddNewUser(uuid string, telegram_id int, email string, limit_ip int, totalGB int, expiryTime int, enable bool, payment bool, start_sub time.Time, end_sub time.Time) bool {
	db, err := connection()
	if err != nil {
		return false
	}
	defer db.Close()

	//query := "INSERT INTO users (uuid, telegram_id, email, limitip, totalgb, expirytime, enable, payment, start_sub, end_sub) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	return false
}
