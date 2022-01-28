package product

import (
	"github.com/DevTeam125/shopping-website/models"
	"gorm.io/gorm"
)

type Feature struct {
	ID          int    `json:"id" gorm:"primary_key"`
	ProductID   int    `json:"article_id"` // Link to Product ID
	Title       string `json:"title"`
	Description string `json:"description"`
}

func SaveProduct(p []Feature) error {

	err := models.DB.Create(p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
