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
			Fracillum string `json:"fracillum"`
			MoonData  []struct {
				Phen string `json:"phen"`
				Time string `json:"time"`
			} `json:"moondata"`
			RiseTime string `json:"riseTime"`
			SetTime  string `json:"setTime"`
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
		// Display error message
		fmt.Println("Error:[API]: Unable to fetch moon data at this time.", err)
	}
	defer resp.Body.Close()

	// Decode JSON response object into MoonData struct
	var moonData MoonData
	err = json.NewDecoder(resp.Body).Decode(&moonData)
	if err != nil {
		// Display error message
		fmt.Println("Error:[JSON decoder]: There was a problem decoding moon data.", err)
	}

	return moonData, nil
}
