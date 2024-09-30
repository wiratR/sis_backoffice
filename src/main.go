package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/wiratR/sis_backoffice/src/api"
	"github.com/wiratR/sis_backoffice/src/database"
)

func main() {
	// Load environment variables from .env file (optional)
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables instead")
	}

	// Initialize the database connection and seed data
	db, err := database.ConnectDB() // This will also seed the images
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	log.Println("Application setup completed successfully!")

	// Create a new Fiber app
	app := fiber.New()

	// Use CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, adjust this as needed for security
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Route to handle image fetching (for both products and services)
	app.Get("/image", api.GetImage)

	// // Start the server on a given port (default 3000 if no PORT env variable)
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "3000"
	// }

	//log.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(":3000"))

}
