package dto

type WeatherOutput struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Current struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float32   `json:"temp_c"`
	TempF            float32   `json:"temp_f"`
	IsDay            int64     `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float32   `json:"wind_mph"`
	WindKph          float32   `json:"wind_kph"`
	WindDegree       int64     `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMB       float32   `json:"pressure_mb"`
	PressureIn       float32   `json:"pressure_in"`
	PrecipMm         float32   `json:"precip_mm"`
	PrecipIn         float32   `json:"precip_in"`
	Humidity         int64     `json:"humidity"`
	Cloud            int64     `json:"cloud"`
	FeelslikeC       float32   `json:"feelslike_c"`
	FeelslikeF       float32   `json:"feelslike_f"`
	VisKM            float32   `json:"vis_km"`
	VisMiles         float32   `json:"vis_miles"`
	Uv               float32   `json:"uv"`
	GustMph          float32   `json:"gust_mph"`
	GustKph          float32   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int64  `json:"code"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float32 `json:"lat"`
	Lon            float32 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}
