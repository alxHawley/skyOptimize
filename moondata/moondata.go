package moondata

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to handle the data from the API
type MoonData struct {
	Geometry struct {
		Coordinates []float64 `json:"coordinates"`
		Type        string    `json:"type"`
	} `json:"geometry"`
	Properties struct {
		Data struct {
			CurPhase  string `json:"curphase"`
			FracIllum string `json:"fracillum"`
			MoonData  []struct {
				Phen string `json:"phen"`
				Time string `json:"time"`
			} `json:"moondata"`
		} `json:"data"`
	} `json:"properties"`
	Type string `json:"type"`
}

// Function to get MoonData struct from API
func GetMoonData(date string, lat string, long string, tz string, dst string) (MoonData, error) {
	url := fmt.Sprintf("https://aa.usno.navy.mil/api/rstt/oneday?date=%s&coords=%s,%s&tz=%s&dst=%s", date, lat, long, tz, dst)
	// Make HTTP request to get JSON response object
	resp, err := http.Get(url)
	if err != nil {
		// Handle error
	}
	defer resp.Body.Close()

	// Decode JSON response object into MoonData struct
	var moonData MoonData
	err = json.NewDecoder(resp.Body).Decode(&moonData)
	if err != nil {
		// Handle error
	}

	// Print out some fields from the MoonData struct
	fmt.Println("Current Phase:", moonData.Properties.Data.CurPhase)
	fmt.Println("Fraction Illuminated:", moonData.Properties.Data.FracIllum)
	fmt.Println("Moon Rise Time:", moonData.Properties.Data.MoonData[0].Time)
	fmt.Println("Moon Set Time:", moonData.Properties.Data.MoonData[2].Time)

	return moonData, nil
}
