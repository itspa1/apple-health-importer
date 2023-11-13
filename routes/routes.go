package routes

import (
	"health-export-parser/controllers"
	"health-export-parser/services"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, healthExportParserService services.HealthExportParser) {
	// initialize health check route
	healthCheckRoutes := router.Group("/health")
	{
		healthCheckRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": "OK",
			})
		})
	}

	// initialize health import routes
	healthImportRoutes := router.Group("/health-import")
	{
		healthImportRoutes.POST("/", func(ctx *gin.Context) {
			controllers.HandleHealthImport(ctx, healthExportParserService)
		})
	}
}
