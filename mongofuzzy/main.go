package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// db.products.createIndex({ description: "text" });
	// ctx := context.Background()
	// MongoDB connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("yourdb").Collection("products")

	// Define the search term
	searchTerm := "your search term"

	// Define the fuzzy matching pipeline
	pipeline := mongo.Pipeline{
		bson.D{{"$search", bson.D{
			{"index", "default"},
			{"text", bson.D{
				{"query", searchTerm},
				{"path", "description"},
				{"fuzzy", bson.D{
					{"maxEdits", 2}, // Adjust the maxEdits value for desired fuzziness
				}},
			}},
		}}},
		bson.D{{"$limit", 10}}, // Limit the number of results
	}

	// Execute the aggregation
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	// Print the results
	for _, result := range results {
		fmt.Println(result)
	}

	// Close the connection
	if err = client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
