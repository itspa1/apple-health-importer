package services

import (
	"context"
	"fmt"
	"health-export-parser/models"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type HealthExportParser interface {
	ParseHeartRateExport(models.HealthData[models.HeartRateMetricExport])
}

type HealthExportParserImpl struct {
	influxBucket   string
	influxWriteApi api.WriteAPIBlocking
}

func NewHealthExportParser(influxWriteApi api.WriteAPIBlocking, influxBucket string) HealthExportParser {
	return &HealthExportParserImpl{
		influxBucket:   influxBucket,
		influxWriteApi: influxWriteApi,
	}
}

func (h HealthExportParserImpl) ParseHeartRateExport(healthData models.HealthData[models.HeartRateMetricExport]) {
	for _, metric := range healthData.Metrics {
		for _, sample := range metric.Data {
			sampleSource := strings.ToLower(sample.Source)
			sampleSource = strings.Join(strings.Split(sampleSource, " "), "_")

			influxMetricName := fmt.Sprintf("%s.%s", metric.Name, sample.Source)

			tags := map[string]string{
				"source": sampleSource,
				"units":  metric.Units,
			}

			fields := map[string]interface{}{
				"max": sample.Max,
				"min": sample.Min,
				"avg": sample.Avg,
			}

			timeParse, err := time.Parse("2006-01-02 15:04:05 -0700", sample.Date)
			if err != nil {
				fmt.Println("error parsing time: ", err)
			}

			newPoint := influxdb2.NewPoint(influxMetricName, tags, fields, timeParse.UTC())

			h.influxWriteApi.WritePoint(context.Background(), newPoint)
		}
	}

	fmt.Println("Successfully wrote to InfluxDB")
}
