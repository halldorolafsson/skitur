package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type skitur struct {
	Dato            string `json:"Dato"`
	Antallkilometer string `json:"Antallkilometer"`
	Antallminutter  string `json:"Antallminutter"`
	Sted            string `json:"Sted"`
}

type alleskiturer []skitur

var skiturer = alleskiturer{
	// {
	// 	Dato:            "",
	// 	Antallkilometer: "",
	// 	Antallminutter:  "",
	// 	Sted:            "",
	// },
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Velkommen til skiturloggen")
}

func createskitur(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var newskitur skitur
	reqBody, err := ioutil.ReadAll(r.Body)
	//Endre feilhåndtering. Denne virker ikke!
	if err != nil {
		fmt.Println(w, "Vennligst legg til dato, antallkilomenter, antallminutter og sted")
	}
	json.Unmarshal(reqBody, &newskitur)
	// Hack to get arount the Cors OPTION stuff that adds a empty entry
	if r.Method == "POST" {
		skiturer = append(skiturer, newskitur)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newskitur)
	}
}

func getAllskiturer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(skiturer)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Content-Type", "application/json")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/skitur", createskitur).Methods("POST", "OPTIONS")
	router.HandleFunc("/skiturer", getAllskiturer).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
