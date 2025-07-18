package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Weather struct {
	Temp          float64 `json:"temp"`
	FeelsLikeTemp float64 `json:"feels_like"`
	Wind          Wind    `json:"wind"`
	Sun           Sun     `json:"sun"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Gust  float64 `json:"gust"`
}

type Sun struct {
	Sunrise int `json:"sunrise"`
	Sunset  int `json:"sunset"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var request Weather
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Fprintln(w, "error unmarshal request: %w", err)
			return
		}

		request.Temp += 10.0

		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(request)
		if err != nil {
			fmt.Fprintln(w, "error marshal data: %w", err)
			return
		}
		w.Write(jsonData)
	}

	if r.Method == http.MethodPost {
		weather := Weather{
			Temp: 12.0,
			Wind: Wind{
				Speed: 5.0,
				Gust:  10.0,
			},
			Sun: Sun{
				Sunrise: 1626076800,
				Sunset:  1626127200,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(weather)
		if err != nil {
			fmt.Fprintln(w, "error marshal data: %w", err)
			return
		}
		w.Write(jsonData)
	}
}

func main() {
	http.HandleFunc("/test", testHandler)
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
