package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"db_connection"
)

type Record struct {
	index      int    `json:"index,omitempty"`
	mac        string `json:"mac,omitempty"`
	cn         string `json:"cn,omitempty"`
	certs      string `json:"cert,omitempty"`
	flash_time string `json:"flash_time,omitempty"`
}

func GetCert(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cn := params["cn"]

	// Call the function to get the certificate
	item, err := db_connection.Get_cert(cn)

	if err != nil {
		// Log the error and return a 404 response
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the raw JSON data directly
	w.Write(item)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/cert/{cn}", GetCert).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
