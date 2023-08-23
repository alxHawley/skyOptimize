package main

import (
	"fmt"
	"html/template"
	"net/http"
	"skyOptimize/geocoder"
	"skyOptimize/moondata"
	moon "skyOptimize/moondata"
	weather "skyOptimize/weatherdata"
)

// Set up localhost server
func main() {
	http.HandleFunc("/stargazer", DataHandler)
	http.ListenAndServe(":8080", nil)
}

type SkyData struct {
	MoonData    []moondata.MoonData `json:"moonData"`
	RiseTime    string              `json:"riseTime"`
	SetTime     string              `json:"setTime"`
	WeatherData weather.WeatherData `json:"weatherData"`
}

// Handler function to get data from APIs and display in HTML template
func DataHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method == "POST" {
		// Parse the form data from the request
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get the form data from the request
		date := r.FormValue("date")
		city := r.FormValue("city")
		state := r.FormValue("state")
		country := r.FormValue("country")
		appid := "f82534b2ff8f9e2c2b0536a99e0d8c87"

		// Get the latitude and longitude for the specified location
		locations, err := geocoder.GetLatLong(fmt.Sprintf("%s,%s,%s", city, state, country), appid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Call moonData function to get MoonData struct for each location
		var moonDataList []moondata.MoonData
		for _, location := range locations {
			moonData, err := moon.GetMoonData(date, fmt.Sprintf("%f", location.Lat), fmt.Sprintf("%f", location.Lon), "", "")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Extract rise and set times from moon data
			var riseTime, setTime string
			for _, data := range moonData.Properties.Data.MoonData {
				if data.Phen == "Rise" {
					riseTime = data.Time
				} else if data.Phen == "Set" {
					setTime = data.Time
				}
			}
			moonData.Properties.Data.RiseTime = riseTime
			moonData.Properties.Data.SetTime = setTime

			moonDataList = append(moonDataList, moonData)
		}

		// Call weatherData function to get WeatherData struct for each location
		var weatherDataList []weather.WeatherData
		for _, location := range locations {
			weatherData, err := weather.GetWeatherData(fmt.Sprintf("%f", location.Lat), fmt.Sprintf("%f", location.Lon), appid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			weatherDataList = append(weatherDataList, weatherData)
		}

		// Combine MoonData and WeatherData into a single struct
		skyData := SkyData{
			MoonData:    moonDataList,
			WeatherData: weatherDataList[0],
		}

		// Create a data map to pass to the stargazer template
		data := map[string]interface{}{
			"MoonData":    skyData.MoonData,
			"WeatherData": skyData.WeatherData,
			"RiseTime":    skyData.MoonData[0].Properties.Data.RiseTime,
			"SetTime":     skyData.MoonData[0].Properties.Data.SetTime,
		}

		// Render the HTML template with the SkyData struct
		stargazerTemplate, err := template.ParseFiles("stargazer.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = stargazerTemplate.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// If the request method is not POST, render the HTML template with the form
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
