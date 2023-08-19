package main

import (
	"html/template"
	"net/http"
	"skyOptimize/moondata"
)

// Set up HTTP server
func main() {
	http.HandleFunc("/moondata", moonDataHandler)
	http.ListenAndServe(":8080", nil)
}

// Handler function for /moondata route
func moonDataHandler(w http.ResponseWriter, r *http.Request) {
	date := "2022-08-16"
	lat := "47.6062"
	long := "-122.3321"
	tz := "-7.0"
	dst := "1"

	// Call moonData function to get MoonData struct
	moonData, err := moondata.GetMoonData(date, lat, long, tz, dst)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse HTML template file
	tmpl, err := template.ParseFiles("skyOptimize.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute HTML template with MoonData struct as data context
	err = tmpl.Execute(w, moonData.Properties.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
