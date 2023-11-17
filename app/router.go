package app

import (
	"fmt"
	"snapfood/config"
	"snapfood/utils/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// initialize router (gin)
func (app *App) initRouter() {
	// init gin
	r := gin.Default()
	// gin middleware config
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		AllowOrigins: []string{
			"http://localhost:9170",
			"http://localhost:9170",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// setup routes
	v1 := r.Group("/v1")
	{
		serviceRoutersV1 := v1.Group("/service")
		serviceRoutersV1.POST("/delay", app.ServiceController.DelayOrder)
		serviceRoutersV1.POST("/assign", app.ServiceController.AssignDelayReport)
		serviceRoutersV1.GET("/reports", app.ServiceController.ReportDelayReport)
	}

	err := r.Run(fmt.Sprintf(":%s", config.Server.HTTPPort))
	if err != nil {
		logger.ZSLogger.Errorf("gin can not be run, err is:%v", err)
	}
}
