package dto

type WeatherOutput struct {
	City  string  `json:"city"`
	TempC float32 `json:"temp_C"`
	TempF float32 `json:"temp_F"`
	TempK float32 `json:"temp_K"`
}
