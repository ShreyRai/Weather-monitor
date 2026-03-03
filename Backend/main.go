package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Weather struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temp"`
	Condition   string  `json:"condition"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS so the frontend can talk to the backend
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "NO_KEY_PROVIDED"
	}

	data := Weather{
		City:        "Mumbai",
		Temperature: 32.5,
		Condition:   "Sunny (Key: " + apiKey + ")",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/api/weather", weatherHandler)
	fmt.Println("Backend starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
