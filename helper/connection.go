package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {

	// Set client options
	clientOptions := options.Client().ApplyURI("cluster_endpoint")
	//Conect to MongoDb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a mongo")
	collection := client.Database("go_rest_api").Collection("books")

	return collection
}

//Error model
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"messge"`
}

//Prepare error model

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())

	var responde = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(responde)

	w.WriteHeader(responde.StatusCode)
	w.Write(message)
}
