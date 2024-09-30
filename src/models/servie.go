package models

import "github.com/google/uuid"

type ImageService struct {
	ID        uuid.UUID `json:"id"` // UUID for the image product
	Index     string    `json:"index"`
	ImageData string    `json:"image_data"` // Base64 encoded image
}
