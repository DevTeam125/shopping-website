package product

import (
	"log"
	"net/http"

	"github.com/DevTeam125/shopping-website/models/product"
	"github.com/gin-gonic/gin"
)

func GetAllProductsBrief(c *gin.Context) {
	var productBrief product.Product
	result, err := productBrief.GetAllProductsBrief(0, 100)
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "result": result})
}

func AddNewProduct(c *gin.Context) {
	var product1 product.Product
	err := c.BindJSON(&product1)

	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	if product1.ID != nil {
		log.Printf("product.ID is set")
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	features := product1.Feature
	photos := product1.Photo

	err = product1.SaveProduct()
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
		return
	}

	if features != nil {
		for i, _ := range features {
			features[i].ProductID = *product1.ID
		}

		err = product.SaveFeatures(features)
		if err != nil {
			log.Printf(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
			return
		}
	}

	if photos != nil {
		for i, _ := range photos {
			photos[i].ProductID = *product1.ID
		}

		err = product.SavePhotos(photos)
		if err != nil {
			log.Printf(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
