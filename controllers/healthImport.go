package controllers

import (
	"fmt"
	"health-export-parser/models"
	"health-export-parser/services"
	"health-export-parser/utils"

	"github.com/gin-gonic/gin"
)

func HandleHealthImport(c *gin.Context, healthExportParserService services.HealthExportParser) {
	typeHeader := c.GetHeader("Record-Type")
	fmt.Println(typeHeader)

	switch typeHeader {
	case "heart-rate":
		var hrHealthExport models.HealthExport[models.HeartRateMetricExport]

		if err := c.BindJSON(&hrHealthExport); err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		utils.SafeExec(func() {
			healthExportParserService.ParseHeartRateExport(hrHealthExport.Data)
		})

	default:
		c.JSON(400, gin.H{
			"error": "Invalid Record-Type",
		})
		return
	}

	c.JSON(200, gin.H{"status": "OK"})

}
