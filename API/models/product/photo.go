package product

import (
	"github.com/DevTeam125/shopping-website/models"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int    `json:"id" gorm:"primary_key"`
	ProductID int    `json:"-"` // Link to Product ID
	Title     string `json:"title"`
	URL       string `json:"url"`
}

func SavePhotos(p []Photo) error {

	err := models.DB.Create(p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
