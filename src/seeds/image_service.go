package seeds

import (
	"database/sql"
	"log"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/wiratR/sis_backoffice/src/models"
	"github.com/wiratR/sis_backoffice/src/utils"
)

// SeedImagesService seeds the image_service table with initial data
func SeedImagesService(db *sql.DB) {
	// get image
	log.Println("start seed images service")

	//Check if the images already exist
	var count int64
	err := db.QueryRow("SELECT COUNT(*) FROM image_service").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to count existing images: %v", err)
	}

	if count > 0 {
		log.Println("Images already seeded, skipping...")
		return
	}
	// Construct the full directory path based on the brand
	dirPath := filepath.Join("image", "service")
	fileCount, fileFullPaths, err := utils.CountFilesInDirectory(dirPath)
	if err != nil {
		log.Fatalf("Error counting files: %v", err)
	}
	log.Printf("Number of files: %d\n", fileCount)
	index := 0
	for _, fullPath := range fileFullPaths {
		index = index + 1
		log.Printf("Index: %d\n", index)
		log.Printf("file name: %s\n", fullPath)
		base64Image, err := utils.GetImageData(fullPath)
		if err != nil {
			log.Fatalf("Error retrieving image data: %v", err)
		}

		image := models.ImageService{
			ID:        uuid.New(),
			Index:     strconv.Itoa(index),
			ImageData: base64Image,
		}

		_, err = db.Exec("INSERT INTO image_service (`id`, `index`, image_data) VALUES (?, ?, ?)", image.ID, image.Index, image.ImageData)
		if err != nil {
			log.Fatalf("Failed to seed image: %v", err)
		}
		log.Printf("Seeded image: %s success\n", fullPath)
	}

	log.Println("Seeding image servcie table completed!")
}
