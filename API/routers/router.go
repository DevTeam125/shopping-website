package routers

import (
	"net/http"
	"os"

	"github.com/DevTeam125/shopping-website/controllers/product"
	"github.com/DevTeam125/shopping-website/middlewares"
	l "github.com/DevTeam125/shopping-website/pkg/logging"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)
	gin.DisableConsoleColor()

	f, err := os.OpenFile("logs/gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		l.Logging.Fatalw("While opening logs/gin.log", "error", err)
	}
	gin.DefaultWriter = f

	router.Use(gin.LoggerWithFormatter(middlewares.Logger()))
	router.Use(gin.CustomRecovery(middlewares.Recovery()))

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
