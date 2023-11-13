package main

import (
	"fmt"
	"health-export-parser/routes"
	"health-export-parser/services"
	"health-export-parser/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Print("Error loading .env file")
		os.Exit(1)
	}

	influxOrg := os.Getenv("INFLUXDB_ORG")
	influxBucket := os.Getenv("INFLUXDB_BUCKET")
	influxToken := os.Getenv("INFLUXDB_TOKEN")
	influxdbURL := os.Getenv("INFLUXDB_URL")

	env := os.Getenv("GIN_ENV")

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	influxWriteApi, err := utils.ConnectInfluxDB(influxdbURL, influxToken, influxOrg, influxBucket)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	healthExportParserService := services.NewHealthExportParser(influxWriteApi, influxBucket)

	r := gin.Default()

	routes.InitializeRoutes(r, healthExportParserService)

	r.Run(":9999")
}
