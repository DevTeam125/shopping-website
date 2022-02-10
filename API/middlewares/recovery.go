package middlewares

import (
	"net/http"

	l "github.com/DevTeam125/shopping-website/pkg/logging"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.RecoveryFunc {
	return func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			l.Logging.Errorw("Gin Panic Recovered", "error", err)
			//c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		l.Logging.Info("Aborted duo to panic recovery")
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
