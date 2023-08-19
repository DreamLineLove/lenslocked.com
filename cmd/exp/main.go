package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
		  id SERIAL PRIMARY KEY,
		  user_id INT NOT NULL,
	    amount INT,
	    description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created!")

	id := 1
	row := db.QueryRow(`
		SELECT name, email FROM users 
		WHERE id=$1;
	`, id)

	var name, email string
	err = row.Scan(&name, &email)
	if errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("User information: name=%q email=%q\n", name, email)

	// user_id := 1
	// for i := 0; i < 5; i++ {
	// 	amount := i * 100
	// 	description := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`
	// 		INSERT INTO orders (user_id, amount, description)
	// 		VALUES ($1, $2, $3);
	// 	`, user_id, amount, description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Sample orders created!")

	user_id := 1
	for i := 6; i < 11; i++ {
		amount := (i - 1) * 100
		description := fmt.Sprintf("Fake order #%d", i-1)
		_, err := db.Exec(`
			UPDATE orders SET amount=$1, description=$2 WHERE id=$3 AND user_id=$4;
		`, amount, description, i, user_id)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Order details modified!")
}
