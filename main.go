package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type skitur struct {
	Dato            string `json:"Dato"`
	Antallkilometer string `json:"Antallkilometer"`
	Antallminutter  string `json:"Antallminutter"`
	Sted            string `json:"Sted"`
}

// cnx is a Connection object to MongoDB to pass around
var cnx = connection()

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Velkommen til skiturloggen")
}

func createSkitur(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var newskitur skitur
	reqBody, err := ioutil.ReadAll(r.Body)
	//Endre feilhåndtering. Denne virker ikke! Hvorfor slår ikke denne til når input er tomt
	if err != nil {
		fmt.Println(w, "Vennligst legg til dato, antallkilomenter, antallminutter og sted")
	}
	json.Unmarshal(reqBody, &newskitur)
	// Hack to get arount the Cors OPTION stuff that adds a empty entry
	if r.Method == "POST" {
		storeSkitur(newskitur)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newskitur)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Content-Type", "application/json")
}
func storeSkitur(s skitur) {
	// Fetch colletion form database using var cnx which is the connection function
	collection := cnx.Database("skilogg").Collection("skitur")
	// Do the insert
	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func connection() *mongo.Client {
	// Set client connection/options // Need to fix this hard coded stuff
	clientOptions := options.Client().ApplyURI("mongodb://skilogg:PWD@USERID.mlab.com:ID/skilogg?retryWrites=false")
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

	return client
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/skitur", createSkitur).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8080", router))
}
