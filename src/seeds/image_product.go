package seeds

import (
	"database/sql"
	"log"
	"path/filepath"
	"strconv"

	// Update this path based on your project structure
	"github.com/google/uuid"
	"github.com/wiratR/sis_backoffice/src/models"
	"github.com/wiratR/sis_backoffice/src/utils"
)

// SeedImagesProduct seeds the image_product table with initial data
func SeedImagesProduct(db *sql.DB) {
	// get image
	log.Println("start seed images product")

	//Check if the images already exist
	var count int64
	err := db.QueryRow("SELECT COUNT(*) FROM image_product").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to count existing images: %v", err)
	}

	if count > 0 {
		log.Println("Images already seeded, skipping...")
		return
	}

	log.Println("list brand name")
	brands, err := utils.ListBrands()
	if err != nil {
		log.Fatalf("Error listing brands: %v", err)
	}
	// Print the list of brands
	for _, brand := range brands {
		// Construct the full directory path based on the brand
		dirPath := filepath.Join("image", "product", brand)
		fileCount, fileFullPaths, err := utils.CountFilesInDirectory(dirPath)
		if err != nil {
			log.Fatalf("Error counting files: %v", err)
		}
		log.Printf("Number of files: %d\n", fileCount)
		index := 0
		for _, fullPath := range fileFullPaths {
			index = index + 1
			log.Printf("Full file paths: %s\n", fullPath)
			log.Printf("Index: %d\n", index)
			base64Image, err := utils.GetImageData(fullPath)
			if err != nil {
				log.Fatalf("Error retrieving image data: %v", err)
			}

			image := models.ImageProduct{
				ID:        uuid.New(),
				Index:     strconv.Itoa(index),
				ImageData: base64Image,
				Brand:     brand,
			}

			_, err = db.Exec("INSERT INTO image_product (`id`, `index`, image_data, brand) VALUES (?, ?, ?, ?)", image.ID, image.Index, image.ImageData, image.Brand)
			if err != nil {
				log.Fatalf("Failed to seed image: %v", err)
			}
			log.Printf("Seeded image: %s success\n", fullPath)
		}

	}

	log.Println("Seeding image product table completed!")
}
