package config

/*
*
* This file is for configuration to connect DB
* * Jangan lupa!!! Untuk penamaan function harus diawali dengan huruf kapital.
 */

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

// let's create a connection to postgre
func CreateConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env") // same as -> var err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// open connection to database
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	// check if connection is error
	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Print("Successfully connected to DB")

	// return the connection
	return db
}

type NullString struct {
	sql.NullString
}

/**
 * kedua fungsi ini berguna untuk menampung jika struct / data bertipe NULL maka dia akan mengisikan data(string) kosong.
 * fungsi ini akan digunakan pada data kosong di models.go
 */
func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
