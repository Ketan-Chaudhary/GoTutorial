package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ketan-chaudhary/mongodb/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//var connectionString = os.Getenv("MONGO_URI")

const dbName = "netflix"
const connectionName = "watchlist"

// Important
var collection *mongo.Collection

// establish connection

func init() { // only run one time when application runs
	//loading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// client option
	clientOption := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	// connect to mongodb
	client, err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection sucess")

	collection = client.Database(dbName).Collection(connectionName)

	// collection
	fmt.Println("Collection reference is ready ")
}

// MongoDb Helper - file

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	out, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted movie in db with id :", out.InsertedID)
}

// Update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Record Updated Successfully with id :", result.UpsertedID)

}

// delete one record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	fmt.Println("Trying to delete ID:", id.Hex())

	filer := bson.M{"_id": id}
	del, err := collection.DeleteOne(context.Background(), filer)

	if err != nil {
		log.Fatal(err)
	}

	if del.DeletedCount == 0 {
		fmt.Println("No document found with that ID.")
	} else {
		fmt.Println("Document deleted successfully.")
	}

	fmt.Println("Deleted Successully ? :", del.Acknowledged)
}

// delete all record
func deleteAllMovie() {
	// filer := bson.D{{}} // select all
	// collection.DeleteMany(context.Background(), filer)
	del, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Records success ? :", del.Acknowledged)
	fmt.Println("No of items deleted :", del.DeletedCount)
}

// read all record
// while reading data we use or mongodb will return a cursor on which we can loop on
// the cursor object have .next() method which will return the next value if exists
// the result will be in format *[]bcon
func getAllMovies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie) // this will not append because the type are mismatched but the primitive are some kind of similar but different is other aspect
		// we just need to use one type consistently
	}

	defer cursor.Close(context.Background())
	return movies
}

// Actual Controller
func GetAllMovies(w http.ResponseWriter, r *http.Request) { // we can use any name
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovie := getAllMovies()
	json.NewEncoder(w).Encode(allMovie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = primitive.NewObjectID()
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methods", "DELETE")

	deleteAllMovie()
	json.NewEncoder(w).Encode("All records Deleted")
}
