package api

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/sis_backoffice/src/database"
	"github.com/wiratR/sis_backoffice/src/models"
)

// fetchImageData retrieves image data and additional info from the database based on the table and index provided
func fetchImageData(table string, index string) (string, string, error) {
	// Connect to the database
	db, err := database.ConnectDB()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return "", "", err
	}
	defer db.Close()

	// Prepare the query dynamically based on the table
	var query string
	var imageData string
	var brand string

	switch table {
	case "image_product":
		var product models.ImageProduct
		query = "SELECT id, `index`, image_data, brand FROM image_product WHERE `index` = ?"
		err = db.QueryRow(query, index).Scan(&product.ID, &product.Index, &product.ImageData, &product.Brand)
		imageData = product.ImageData
		brand = product.Brand
	case "image_service":
		var service models.ImageService
		query = "SELECT id, `index`, image_data FROM image_service WHERE `index` = ?"
		err = db.QueryRow(query, index).Scan(&service.ID, &service.Index, &service.ImageData)
		imageData = service.ImageData
		brand = "" // No brand for services
	default:
		return "", "", fiber.NewError(fiber.StatusInternalServerError, "Unknown table")
	}

	if err != nil {
		log.Println("Item not found or error fetching data:", err)
		return "", "", fiber.NewError(fiber.StatusNotFound, "Item not found")
	}

	// Extract only the Base64 part if the data contains metadata
	if strings.Contains(imageData, ",") {
		parts := strings.Split(imageData, ",")
		if len(parts) > 1 {
			imageData = parts[1]
		}
	}

	return imageData, brand, nil
}

// // serveImage decodes Base64 and sends the image with appropriate content type
// func serveImage(c *fiber.Ctx, imageData string) error {
// 	// Decode the Base64 string
// 	imageBytes, err := base64.StdEncoding.DecodeString(imageData)
// 	if err != nil {
// 		log.Println("Error decoding Base64 image:", err)
// 		return c.Status(fiber.StatusInternalServerError).SendString("Error decoding image")
// 	}

// 	// Set the content type dynamically if necessary (currently assuming PNG)
// 	c.Set("Content-Type", "image/png")

// 	// Send the decoded image bytes as the response
// 	return c.Send(imageBytes)
// }

// GetImage handles both product and service image requests and returns JSON response
func GetImage(c *fiber.Ctx) error {
	// Determine if the query is for a product or a service
	productIndex := c.Query("product")
	serviceIndex := c.Query("service")

	var table string
	var index string

	if productIndex != "" {
		table = "image_product"
		index = productIndex
	} else if serviceIndex != "" {
		table = "image_service"
		index = serviceIndex
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameter. Please provide 'product' or 'service'.",
		})
	}

	imageData, brand, err := fetchImageData(table, index)
	if err != nil {
		return err
	}

	// Return the image data and additional information as a JSON response
	response := fiber.Map{
		"index":      index,
		"image_data": imageData, // Base64 image data
	}

	if table == "image_product" {
		response["brand"] = brand
	}

	return c.JSON(response)
	//return serveImage(c, imageData)
}
