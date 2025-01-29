package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// TODO: Set up to listen at specific domain
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Retrieve query values needed for data call
	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, "Error parsing query parameters", http.StatusBadRequest)
		return
	}

	// Throw error if not at root
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// Throw error if not get
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve required params
	latitude := queryValues.Get("latitude")
	longitude := queryValues.Get("longitude")
	timezone := queryValues.Get("timezone")
	days := queryValues.Get("days")

	if latitude == "" || longitude == "" || timezone == "" || days == "" {
		http.Error(w, "lattitude, longitude, timezone, and days are required", http.StatusNotAcceptable)
		return
	}

	// Combine params into weather request
	weatherUrl := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=temperature_2m,relative_humidity_2m,precipitation_probability,wind_speed_10m&temperature_unit=fahrenheit&wind_speed_unit=mph&precipitation_unit=inch&timezone=%s&forecast_days=%s", latitude, longitude, timezone, days)
	resp, err := http.Get(weatherUrl)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	defer resp.Body.Close()

	// Ready body of email and throw error if necessary
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var forcastData ForcastData
	err = json.Unmarshal(body, &forcastData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Convert to ingestible format for D3.js chart locally
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(forcastData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
