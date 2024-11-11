package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func processReceipt(w http.ResponseWriter, r *http.Request) {

	// Read the payload from the request.
	payload, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	var receipt Receipt

	// Parse the json payload to the receipt type.
	err = json.Unmarshal(payload, &receipt)

	if err != nil {
		panic(err)
	}

	// Initialize and validate payload.
	validate := validator.New()

	err = validate.Struct(receipt)

	if err != nil {
		panic(err)
	}

	// Compute and get points from the receipt object.
	var points = calculatePoints(receipt)

	// Build the response object with the uuid.
	id := uuid.New()
	response := ProcessResponse{
		Id: id.String(),
	}

	// Store the id and points as key-value pair in processData map/dictionary.
	processData[id.String()] = points

	// Encode as json and return the response with Content-Type header.
	jsonData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handleReceiptPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	if point, ok := processData[id]; ok {

		response := PointsResponse{
			Points: point,
		}

		data, _ := json.Marshal(response)

		w.Write(data)
	} else {
		invalid := map[string]interface{}{
			"message": "Invalid Id",
			"error":   true,
		}

		data, _ := json.Marshal(invalid)
		w.Write(data)
	}
}
