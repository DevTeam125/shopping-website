package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DevTeam125/shopping-website/models/product"
	"github.com/gin-gonic/gin"
)

func GetProductByID(c *gin.Context) {
	var product product.Product
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	result, err := product.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "result": result})
}

func GetAllProductsBrief(c *gin.Context) {
	var productBrief product.Product
	result, err := productBrief.GetAllProductsBrief(0, 100)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "result": result})
}

func AddNewProduct(c *gin.Context) {
	var product1 product.Product
	err := c.BindJSON(&product1)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	if product1.ID != nil {
		log.Println("product.ID is set")
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	if product1.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Field name is mandatory"})
		return
	}

	features := product1.Feature
	photos := product1.Photo

	err = product1.SaveProduct()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
		return
	}

	if features != nil {
		for i := range features {
			features[i].ProductID = *product1.ID
		}

		err = product.SaveFeatures(features)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
			return
		}
	}

	if photos != nil {
		for i := range photos {
			photos[i].ProductID = *product1.ID
		}

		err = product.SavePhotos(photos)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func UpdateProduct(c *gin.Context) {
	var productToBeUpdated product.Product
	err := c.BindJSON(&productToBeUpdated)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	if productToBeUpdated.ID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Field ID is mandatory"})
		return
	}

	done, err := productToBeUpdated.UpdateProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	if done {
		c.JSON(http.StatusOK, gin.H{"ok": true})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"ok": false})
}

func DeleteProductByID(c *gin.Context) {
	var product product.Product

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	done, err := product.DeleteProductByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	if !done {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
