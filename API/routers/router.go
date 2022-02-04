package routers

import (
	"net/http"

	"github.com/DevTeam125/shopping-website/controllers/product"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	productGroup := router.Group("/product")
	{
		productGroup.GET("", product.GetAllProductsBrief)
		productGroup.GET(":id", product.GetProductByID)
		productGroup.POST("", product.AddNewProduct)
		productGroup.PUT("", product.UpdateProduct)
		productGroup.DELETE(":id", product.DeleteProductByID)
	}

	return router
}
