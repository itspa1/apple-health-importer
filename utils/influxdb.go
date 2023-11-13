package utils

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func ConnectInfluxDB(url string, token string, org string, bucket string) (api.WriteAPIBlocking, error) {
	client := influxdb2.NewClient(url, token)

	client.Options().WriteOptions().SetBatchSize(500) // write every 500 points

	_, err := client.Health(context.Background())
	if err != nil {
		return nil, err
	}

	writeApiBlockingApi := client.WriteAPIBlocking(org, bucket)

	return writeApiBlockingApi, nil
}
