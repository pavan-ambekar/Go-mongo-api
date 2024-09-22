package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pavan-ambekar/mongo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string

func init() {
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	connectionString = fmt.Sprintf("mongodb+srv://%s:%s@cluster0.s53nauh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", mongoUser, mongoPassword)
}

const dbName = "movies"
const colName = "watchList"

var collection *mongo.Collection

// connect with mongoDB
func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Mongo connection success")

	collection = client.Database(dbName).Collection(colName)
}

// Mongodb helpers

// insert one record
func insertOneMovie(movie models.Movie) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 Movie with id:", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted count:", result.DeletedCount)
}
