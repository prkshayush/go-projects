package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mongo/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const connectionString = "mongodb+srv://prkshayush:alwaysAyush@database-golang.fpqf0ye.mongodb.net/?retryWrites=true&w=majority"
const dbName = "DigiBookLib"
const colName = "book-collection"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

func insertOneBook(book model.Library) {
	inserted, err := collection.InsertOne(context.Background(), book)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Inserted one book name in database with id: ", inserted.InsertedID)
}

func updateOneBook(bookID string) {
	id, _ := primitive.ObjectIDFromHex(bookID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"Read": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified coount: ", result.ModifiedCount)
}

func deleteOneBook(bookID string) {
	id, _ := primitive.ObjectIDFromHex(bookID)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted book count: ", deleteCount)
}

func deleteAllBook() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Books deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllBooks() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var books []primitive.M

	for curr.Next(context.Background()) {
		var book bson.D
		err := curr.Decode(&book)

		if err != nil {
			log.Fatal(err)
		}
		books = append(books)

	}
	defer curr.Close(context.Background())
	return books
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allBooks := getAllBooks()
	json.NewEncoder(w).Encode(allBooks)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var book model.Library
	_ = json.NewDecoder(r.Body).Decode(&book)
	insertOneBook(book)
	json.NewEncoder(w).Encode(book)
}

func MarkAsRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneBook(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteABook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneBook(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllBook()
	json.NewEncoder(w).Encode(count)
}
