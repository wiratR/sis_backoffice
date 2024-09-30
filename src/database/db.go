package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver for SQL
	"github.com/joho/godotenv"         // Import godotenv for loading .env files
	"github.com/wiratR/sis_backoffice/src/seeds"
)

// ConnectDB establishes a connection to the MariaDB database
// Returns a pointer to the database and any connection error encountered
func ConnectDB() (*sql.DB, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		return nil, err
	}

	// Read database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Data Source Name (DSN) for connecting to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	// dsn := "polar_admin:Sispass.2024@!@tcp(147.50.231.19:3306)/polar_sis"

	// sql.Open initializes a database object for the DSN but doesn't verify the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// If there's an issue with the DSN or initializing the connection
		log.Printf("Error opening database: %v\n", err)
		return nil, err
	}

	// Ping verifies the actual connection to the database
	err = db.Ping()
	if err != nil {
		// If the database isn't reachable, this will return an error
		log.Printf("Error connecting to database: %v\n", err)
		return nil, err
	}

	fmt.Println("Successfully connected to MariaDB!")

	// Seed the database with images
	seeds.SeedImagesProduct(db)
	// Seed the database with images
	seeds.SeedImagesService(db)

	return db, nil
}
