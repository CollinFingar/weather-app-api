package main

type ForcastData struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone string `json:"timezone"`
	Units HourlyUnits `json:"hourly_units"`
	Hourly HourlyData `json:"hourly"`
}

type HourlyUnits struct {
    Time string `json:"time"`
	Temperature string `json:"temperature_2m"`
	Humidity string `json:"relative_humidity_2m"`
	PrecipitationProbability string `json:"precipitation_probability"`
	WindSpeed string `json:"wind_speed_10m"`
}

type HourlyData struct {
	Time []string `json:"time"`
	Temperature []float64 `json:"temperature_2m"`
	Humidity []int64 `json:"relative_humidity_2m"`
	PrecipitationProbability []int64 `json:"precipitation_probability"`
	WindSpeed []float64 `json:"wind_speed_10m"`
}
