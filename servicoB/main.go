package main

import (
	"encoding/json"
	"net/http"
)

type WeatherResponse struct {
    City   string  `json:"city"`
    TempC  float64 `json:"temp_C"`
    TempF  float64 `json:"temp_F"`
    TempK  float64 `json:"temp_K"`
}

func main() {
    http.HandleFunc("/get-weather", GetWeather)
    http.ListenAndServe(":8081", nil)
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
    // Here you should receive the CEP from Service A
    // Code for receiving CEP goes here

    // Here you should call the external API (like viaCEP) to get location data
    // Code for calling external API goes here

    // Here you should call the external API (like WeatherAPI) to get weather data
    // Code for calling external API goes here

    // Here you should format the response
    // Code for formatting response goes here

    // Return response
    weatherResponse := WeatherResponse{
        City:  "São Paulo",
        TempC: 28.5,
        TempF: 28.5 * 1.8 + 32,
        TempK: 28.5 + 273,
    }
    jsonResponse, err := json.Marshal(weatherResponse)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
