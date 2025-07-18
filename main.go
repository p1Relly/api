package main

import (
	"encoding/json"
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
	if r.Method == http.MethodPost {
		weather := Weather{
			Temp: 36.6,
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}

	if r.Method == http.MethodGet {
		weather := Weather{
			Temp: 12.0,
		}

		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(weather)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
