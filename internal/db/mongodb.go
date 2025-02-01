package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// For docker compose
const MONGODB_URI = "mongodb://root:password@db:27017"

// For local development
// const MONGODB_URI = "mongodb://root:password@localhost:27017"

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

func GetMetrics(client *mongo.Client, skip int64, limit int64) ([]Metric, error) {
	filter := bson.D{}
	sort := bson.D{{Key: START_TIME_FIELD, Value: -1}}
	opts := options.Find().SetSkip(skip).SetLimit(limit).SetSort(sort)
	cursor, err := client.Database(DB).Collection(COLLECTION).Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}

	var metrics []Metric
	if err = cursor.All(context.Background(), &metrics); err != nil {
		return nil, err
	}

	return metrics, nil
}

func DeleteMetric(client *mongo.Client, id string) (bool, error) {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("error converting id to objectId: %v\n", err)
	}

	filter := bson.D{{Key: ID_FIELD, Value: objectId}}
	res, err := client.Database(DB).Collection(COLLECTION).DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if res.DeletedCount == 0 {
		return false, nil
	}
	return true, nil
}

func DocumentCount(client *mongo.Client) (int64, error) {
	res, err := client.Database(DB).Collection(COLLECTION).CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return -1, err
	}

	return res, nil
}
