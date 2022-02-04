package routers

import (
	"net/http"
	"time"

	"github.com/DevTeam125/shopping-website/controllers/product"
	l "github.com/DevTeam125/shopping-website/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRoutes() *gin.Engine {
	router := gin.New()
	logger := l.ZapLogger
	defer logger.Sync()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		logger.Info("",
			zap.String("Method", param.Method),
			zap.String("IP", param.ClientIP),
			zap.String("TS", param.TimeStamp.Format(time.RFC1123)),
			zap.String("Path", param.Path),
			zap.String("Proto", param.Request.Proto),
			zap.Int("StatusCode", param.StatusCode),
			zap.Duration("Latency", param.Latency),
			zap.String("UserAgent", param.Request.UserAgent()),
			zap.String("ErrorMessage", param.ErrorMessage),
		)

		return ""
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
