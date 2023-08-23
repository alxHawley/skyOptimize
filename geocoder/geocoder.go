package geocoder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	// Country string  `json:"country"`
	State string `json:"state"`
}

// Direct geocode function to get Lat/Long struct from API
func GetLatLong(q string, appid string) ([]Location, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", q, appid)
	// Make HTTP request to get JSON response object
	resp, err := http.Get(url)
	if err != nil {
		// Display error message
		fmt.Println("Error:[API]: Unable to geocode location data at this time.", err)
	}
	defer resp.Body.Close()

	// Decode JSON response object into slice of Location structs
	var locations []Location
	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		// Display error message
		fmt.Println("Error:[JSON decoder]: There was a problem decoding location data.", err)
	}

	return locations, nil
}
