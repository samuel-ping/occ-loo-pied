package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const MONGODB_URI = "mongodb://root:password@db:27017"
const DB = "occloopied"
const COLLECTION = "metrics"

func ConnectMongo() (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func AddOccupiedMetric(client *mongo.Client, startTime *time.Time, endTime *time.Time) error {
	duration := endTime.Sub(*startTime)

	_, err := client.Database(DB).Collection(COLLECTION).InsertOne(
		context.Background(),
		map[string]interface{}{
			"startTime": startTime,
			"endTime":   endTime,
			"duration":  duration,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
