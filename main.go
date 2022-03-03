package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go/api/helper"
	"go/api/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()
var collection1 = helper.ConnectDB1()

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var comments []models.Comment

	// bson.M{},
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// single document can be decoded
		var cmt models.Comment

		err := cur.Decode(&cmt)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		comments = append(comments, cmt)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(comments) // encode similar to serialize process.
}

func getOneComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var cmt models.Comment

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&cmt)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(cmt)

}

func createComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var comment models.Comment

	_ = json.NewDecoder(r.Body).Decode(&comment)

	result, err := collection.InsertOne(context.TODO(), comment)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

//create reply
func createReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reply models.Reply

	//to get comment

	_ = json.NewDecoder(r.Body).Decode(&reply)

	// insert  Reply model.
	result, err := collection1.InsertOne(context.TODO(), reply)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

//update reply

// func updateReplay(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var params = mux.Vars(r)

// 	//Get id from parameters
// 	id, _ := primitive.ObjectIDFromHex(params["id"])

// 	var cmt models.Comment

// 	// Create filter
// 	filter := bson.M{"_id": id}

// 	// Read update model from body request
// 	_ = json.NewDecoder(r.Body).Decode(&cmt)

// 	// prepare update model.
// 	s
// 	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&cmt)

// 	if err != nil {
// 		helper.GetError(err, w)
// 		return
// 	}

// 	cmt.ID = id

// 	json.NewEncoder(w).Encode(cmt)
// }

// var client *mongo.Client

func main() {

	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/comments", getComments).Methods("GET")
	r.HandleFunc("/api/comment/{id}", getOneComment).Methods("GET")
	r.HandleFunc("/api/createcomment", createComment).Methods("POST")
	r.HandleFunc("/api/replycmt/{id}", createReply).Methods("POST")
	//r.HandleFunc("/api/cmtwithchat", getCmtWithChat).Methods("GET")

	//r.HandleFunc("/api/reply/{id}", updateReplay).Methods("PUT")

	//r.HandleFunc("/api/books/{id}", deleteComment).Methods("DELETE")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))
	// if err := http.ListenAndServe(":80", nil); err != nil {
	// 	panic(err)
	// }

}
