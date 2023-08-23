package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to handle the data from the API
type WeatherData struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp    float64 `json:"temp"`
		TempMin float64 `json:"temp_min"`
		TempMax float64 `json:"temp_max"`
	} `json:"main"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Name string `json:"name"`
}

// Function to get WeatherData struct from API
func GetWeatherData(lat string, long string, appid string) (WeatherData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=imperial", lat, long, appid)
	// Make HTTP request to get JSON response object
	resp, err := http.Get(url)
	if err != nil {
		// Display error message
		fmt.Println("Error:[API]: Unable to fetch weather data at this time.", err)
	}
	defer resp.Body.Close()

	// Decode JSON response object into WeatherData struct
	var weatherData WeatherData
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		// Display error message
		fmt.Println("Error:[JSON decoder]: There was a problem decoding weather data.", err)
	}

	return weatherData, nil
}
