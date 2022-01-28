package router

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

	articleGroup := router.Group("/product")
	{
		articleGroup.GET("", product.GetAllProductsBrief)
		articleGroup.POST("", product.AddNewProduct)
	}

	return router
}
