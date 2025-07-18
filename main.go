package main

import (
	"encoding/json"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	weather := Weather{
		Temp: 36.6,
	}

	if r.Method == http.MethodGet {
		//
	}

	// Encode the response as JSON
	jsonData, err := json.Marshal(weather)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the response writer
	w.Write(jsonData)
}

type Weather struct {
	Temp          float64 `json:"temp"`
	FeelsLikeTemp float64 `json:"feels_like"`
	Wind          Wind
	Sun           Sun
}

type Wind struct {
	Speed float64 `json:"speed"`
	Gust  float64 `json:"gust"`
}

type Sun struct {
	Sunrise int `json:"sunrise"`
	Sunset  int `json:"sunset"`
}

func main() {
	http.HandleFunc("/", testHandler)

	http.ListenAndServe("5.189.237.243:8082", nil)
}
