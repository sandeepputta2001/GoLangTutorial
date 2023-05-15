package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://puttasandeep2001:<password>@cluster0.ucyhumg.mongodb.net/?retryWrites=true&w=majority"
const colName = "Watchlist"
const dbName = "Netflix"

// Most important : here we are declaring the reference of the collection variable because it reduces the effort of searching for the collection digging all the existing variables
var collection *mongo.Collection

// connect to mongodb

func init() {
	clientOption := options.Client().ApplyURI(connectionString) //asking client to choose component to connect with mongo db

	client, err := mongo.Connect(context.TODO(), clientOption) // using mongo package/driver connect the client to mongodb with the selected clientoption and this client is now capable of connecting with database

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb is connected successfully")
	collection = client.Database(dbName).Collection(colName) // using the client connection with the mongo we have created the database called dbName and collecion called colName
}

func insertOneMovie(movie model.Netflix) {
	//mongodb helper method : they insert data into the database

	// insert 1 record

	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the id of the inserted value is ", inserted.InsertedID)
}

func updateOneMovie(movieID string) {

	id, _ := primitive.ObjectedFromHex                //object from hex converts string datatype of id int objectid
	filter := bson.M{"_id": id}                       // db always concerned with bson package bson.M{}(used to pass parameters ) and through this line we are filtering the record or data or collection to be updated
	update := bson.M{"$set": bson.M{"watched": true}} // here we are finding the update variable to be updated using the $ set operator

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count :", result.ModifiedCount)
}

// delete one record

func deleteOneRecord(movieID string) {
	id._ := primitive.ObjectedFromHex
	filter := bson.M{"_id": id}
	deleteCount := collecion.deleteOne(context.Background(), filter) // deleteOne method returns deleted count

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("movie got deleted with delete count :", deleteCount)

}

// delete all values
func deleteAll(movie string) int64 { // deletedcount is of int64 type
	deletedResult, err := collecion.DeleteMany(context.Background(), bson.D{{}}, nil) // empty set deletes all the values and
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of movies delete :", deletedResult.DeletedCount) //deletecount in deleteMany consists of key value pairs
	return deleteAll.DeletedCount
}

// get all movies/records from database

func getAllMovies() {
	cursor, err := collecion.Find(ctx.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(ctx.Background()) {

		var movie bson.M

		if err := cur.Decode(&movie); err != nil {

			log.Fatal(err)
		}
		movies = append(movies, movie)

	}

	defer cur.Close(context.Background())
	return movies

}

// Actual controller    -file(file indicates that controllers should write in the seperate page )

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

// inserting a movie into mongodb

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-allow-origin", "POSTS")

	var movie model.Netflix

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

// marking movie as watched

func MarkAsWathced(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-allow-methods", "PUT")

	params := mux.vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

// delete a movie

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-allow-methods", "DELETE")

	params := mux.Vars(r)

	deleteOneRecord(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-allow-methods", "DELETE")

	count := deleteAll() // delete all method returns count
	json.NewEncoder(w).Encode(count)
}
