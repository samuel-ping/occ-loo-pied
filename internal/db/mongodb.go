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

func UsagesByDay(client *mongo.Client) ([]UsagesByDayMetric, error) {
	groupStage := bson.D{
		{Key: GROUP, Value: bson.D{
			{Key: ID_FIELD, Value: bson.D{
				{Key: DATE_TO_STRING, Value: bson.D{
					{Key: FORMAT, Value: "%Y-%m-%d"},
					{Key: DATE, Value: "$" + START_TIME_FIELD},
				}},
			}},
			{Key: COUNT_FIELD, Value: bson.D{
				{Key: SUM, Value: 1},
			}},
		}},
	}

	cursor, err := client.Database(DB).Collection(COLLECTION).Aggregate(context.Background(), mongo.Pipeline{groupStage})
	if err != nil {
		return nil, err
	}

	var usagesByDay []UsagesByDayMetric
	if err = cursor.All(context.Background(), &usagesByDay); err != nil {
		return nil, err
	}

	return usagesByDay, nil
}

func CalcStats(client *mongo.Client) (Stats, error) {
	var stats Stats

	totalUsages, err := DocumentCount(client)
	if err != nil {
		return stats, err
	}
	stats.TotalUsages = totalUsages

	facetStage := bson.D{
		{Key: FACET, Value: bson.D{
			{Key: TOTAL_DURATION_FIELD, Value: totalDurationPipeline()},
			{Key: LONGEST_DURATION_AND_DATE_FIELD, Value: longestDurationAndDatePipeline()},
			{Key: AVERAGE_DURATION_FIELD, Value: avgDurationPipeline()},
		}},
	}

	// moves all the important fields to the top level of the document
	projectStage := bson.D{
		{Key: PROJECT, Value: bson.D{
			{Key: TOTAL_DURATION_FIELD, Value: bson.D{
				{Key: ARRAY_ELEM_AT, Value: bson.A{"$" + TOTAL_DURATION_FIELD + "." + TOTAL_DURATION_FIELD, 0}},
			}},
			{Key: LONGEST_DURATION_AND_DATE_FIELD, Value: bson.D{
				{Key: ARRAY_ELEM_AT, Value: bson.A{"$" + LONGEST_DURATION_AND_DATE_FIELD, 0}},
			}},
			{Key: AVERAGE_DURATION_FIELD, Value: bson.D{
				{Key: ARRAY_ELEM_AT, Value: bson.A{"$" + AVERAGE_DURATION_FIELD + "." + AVERAGE_DURATION_FIELD, 0}},
			}},
		}},
	}

	cursor, err := client.Database(DB).Collection(COLLECTION).Aggregate(context.Background(), mongo.Pipeline{facetStage, projectStage})
	if err != nil {
		return stats, err
	}

	var durationStats []DurationStats
	if err = cursor.All(context.Background(), &durationStats); err != nil {
		return stats, err
	}
	stats.DurationStats = durationStats[0]

	return stats, nil
}

func totalDurationPipeline() bson.A {
	groupStage := bson.D{
		{Key: GROUP, Value: bson.D{
			{Key: ID_FIELD, Value: nil},
			{Key: TOTAL_DURATION_FIELD, Value: bson.D{
				{Key: SUM, Value: "$" + DURATION_FIELD},
			}},
		}},
	}

	return bson.A{groupStage}
}

func longestDurationAndDatePipeline() bson.A {
	sortStage := bson.D{{Key: SORT, Value: bson.D{{Key: DURATION_FIELD, Value: -1}}}}
	limitStage := bson.D{{Key: LIMIT, Value: 1}}

	return bson.A{sortStage, limitStage}

}

func avgDurationPipeline() bson.A {
	groupStage := bson.D{
		{Key: GROUP, Value: bson.D{
			{Key: ID_FIELD, Value: nil},
			{Key: AVERAGE_DURATION_FIELD, Value: bson.D{
				{Key: AVG, Value: "$" + DURATION_FIELD},
			}},
		}},
	}

	return bson.A{groupStage}
}
