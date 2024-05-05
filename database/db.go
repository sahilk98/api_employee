package database

import (

	"database/sql"
	"os"
	"fmt"

	_ "github.com/lib/pq" // Import postgres driver

)

func ConnectDB()(*sql.DB, error){

	// not the best way, but its only a local db
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "admin"
	dbname := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to DB: ", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error while trying to Ping DB: ", err)
	}

	return db, nil
}

func MigrateTable() error {

	dbConnection, errConnect := ConnectDB()
	if errConnect != nil {
		return fmt.Errorf("error while connecting to DB: ", errConnect)
	}

	defer dbConnection.Close()

	query, errReadFile := ReadFile("database/migrations/create_table.up.sql")
	if errReadFile != nil {
		return fmt.Errorf("error while reading migration file: ", errReadFile)
	}

	_, errExecQuery := dbConnection.Exec(query)
	if errExecQuery != nil{
		return fmt.Errorf("error while executing query: ", errExecQuery)
	}

	return nil
}

func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(data), nil
}