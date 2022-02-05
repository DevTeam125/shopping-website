package routers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DevTeam125/shopping-website/controllers/product"
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

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)

	}))
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			l.Logging.Errorw("Gin Panic Recovered", "error", err)
			//c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		l.Logging.Info("Aborted duo to panic recovery")
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	router.GET("/", func(c *gin.Context) {
		//panic("Got Here")
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
